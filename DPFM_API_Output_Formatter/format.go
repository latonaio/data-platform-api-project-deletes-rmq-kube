package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToProject(rows *sql.Rows) (*Project, error) {
	defer rows.Close()
	project := Project{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&project.project,
			&project.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &project, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &project, nil
	}

	return &project, nil
}
