package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-project-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-project-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) Project(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Project {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE project.Project = %d ", input.Project.Project),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	project.Project
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_project_data as project 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProject(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
