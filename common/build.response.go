package common

//DefaultResponse default payload response
type DefaultResponse struct {
	ResultCode string `json:"resultCode"`
	HttpStatus string `json:"http_status"`
	Message    string `json:"developerMessage"`
}

func ResponseOk(err string) DefaultResponse {
	return DefaultResponse{
		"20000",
		"200",
		"success",
	}
}

func ResponseOkWithCondition(err string) DefaultResponse {
	return DefaultResponse{
		"20001",
		"200",
		"success_with_condition",
	}
}

func ResponseOkDataNotFound(err string) DefaultResponse {
	return DefaultResponse{
		"20020",
		"200",
		"data_not_found",
	}
}

func ResponseOkDataCreatedSuccess(err string) DefaultResponse {
	return DefaultResponse{
		"20100",
		"201",
		"created_success",
	}
}

func ResponseBadRequest(err string) DefaultResponse {
	return DefaultResponse{
		"40000",
		"400",
		"bad_request",
	}
}

func ResponseUnAuthorized(err string) DefaultResponse {
	return DefaultResponse{
		"40101",
		"401",
		"unauthorized",
	}
}

func ResponseForm1Forbidden(err string) DefaultResponse {
	return DefaultResponse{
		"40300",
		"403",
		"missing_or_invalid_parameter" + err,
	}
}

func ResponseDataExist(err string) DefaultResponse {
	return DefaultResponse{
		"40301",
		"403",
		"data_exist",
	}
}

func ResponseQuotaExceded(err string) DefaultResponse {
	return DefaultResponse{
		"42200",
		"422",
		"quota_exceded",
	}
}

func ResponsDataNotFound(err string) DefaultResponse {
	return DefaultResponse{
		"40401",
		"404",
		"created_success",
	}
}

func ResponsIse1SystemError(err string) DefaultResponse {
	return DefaultResponse{
		"50000",
		"500",
		"system_error",
	}
}

func ResponsIse2SystemError(err string) DefaultResponse {
	return DefaultResponse{
		"50001",
		"500",
		"created_success",
	}
}

func ResponsConnectionTimeOut(err string) DefaultResponse {
	return DefaultResponse{
		"5002",
		"500",
		"connection_timeout",
	}
}

func ResponsConnectionError(err string) DefaultResponse {
	return DefaultResponse{
		"50003",
		"500",
		"connection_error",
	}
}

func ResponsIseQueryError(err string) DefaultResponse {
	return DefaultResponse{
		"50004",
		"500",
		"execute_query_error",
	}
}

func ResponsInsertError(err string) DefaultResponse {
	return DefaultResponse{
		"50005",
		"500",
		"execute_insert_error",
	}
}

func ResponsUpdateError(err string) DefaultResponse {
	return DefaultResponse{
		"50006",
		"500",
		"execute_update_error",
	}
}

func ResponsDeleteError(err string) DefaultResponse {
	return DefaultResponse{
		"50007",
		"500",
		"execute_delete_error",
	}
}

func ResponsIse8UnknowError(err string) DefaultResponse {
	return DefaultResponse{
		"50060",
		"500",
		"unknow_error",
	}
}

func ResponsSu1ServerBusy(err string) DefaultResponse {
	return DefaultResponse{
		"50300",
		"503",
		"server_busy",
	}
}

func ResponsServerUnvailable(err string) DefaultResponse {
	return DefaultResponse{
		"50301",
		"503",
		"server_unavailable",
	}
}

func ResponsGateWayTimeout(err string) DefaultResponse {
	return DefaultResponse{
		"50400",
		"504",
		"gateway_timeout_error",
	}
}
