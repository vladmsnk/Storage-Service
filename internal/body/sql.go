package body

const (
	SQLInsertMaterialInfo = `INSERT INTO public.material_info 
		(material_id, author, title, file_type, file_link) 
		VALUES (:material_id, :author, :title, :file_type, :file_link);`
)
