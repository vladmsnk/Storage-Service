//go:build integration

package environment

import (
	"github.com/ory/dockertest/v3"
	"log"
	"os/exec"

	"github.com/ory/dockertest/v3/docker"
)

var Env Environment

type Environment struct {
	Pool     *dockertest.Pool
	Postgres *dockertest.Resource
	Redis    *dockertest.Resource
	RMQ      *dockertest.Resource
}

func StartPool() error {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Printf("Не получилось подключиться к докеру: %s", err)
		return err
	}
	Env.Pool = pool
	return nil
}

func StartContainer(pool *dockertest.Pool, options *dockertest.RunOptions) (*dockertest.Resource, error) {
	resource, ok := pool.ContainerByName(options.Name)

	if ok {
		err := StopContainer(resource)
		if err != nil {
			log.Printf("Не получилось перезапустить контейнер: %s", err)
			return nil, err
		}
	}

	resource, err := pool.RunWithOptions(options, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Printf("Не получилось запустить контейнер: %s", err)
		return nil, err
	}

	return resource, nil
}

func StopContainer(resource *dockertest.Resource) error {
	if err := resource.Close(); err != nil {
		log.Printf("Не получилось остановить контейнер %s", err)
		return err
	}
	return nil
}

func DumpContainerLogs(resource *dockertest.Resource) error {
	out, err := exec.Command("docker", "logs", resource.Container.ID).CombinedOutput() //#nosec G204 Линтер предупреждает,
	if err != nil {
		log.Printf("Не получилось получить логи контейнера: %v", err)
		return err
	}
	log.Printf("Logs for %s\n%s:", resource.Container.ID, out)
	return nil
}
