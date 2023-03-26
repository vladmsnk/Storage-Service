package body

const (
	SQLInsertMaterialInfo = `INSERT INTO public.material_info 
		(material_id, author, title, file_type, file_link) 
		VALUES ($1, $2, $3, $4, $5);`

	SQLDeleteMaterialInfo = `DELETE FROM public.material_info WHERE material_id = $1;`
)
