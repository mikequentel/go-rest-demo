package model

import (
	"encoding/json"
	"time"
  uuid "github.com/satori/go.uuid"
)

type Restaurant struct {
	ID                            uuid.UUID     `json:"id" db:"id"`
	OID                           int           `json:"oid" db:"oid"`
	CreatedAt                     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt                     time.Time    `json:"updated_at" db:"updated_at"`
	Address                       string  `json:"address" db:"address"`
	County                        string  `json:"county" db:"county"`
	CriticalViolation             string  `json:"critical_violation" db:"critical_violation"`
	DateOfInspection              time.Time    `json:"date_of_inspection" db:"date_of_inspection"`
	Facility                      string  `json:"facility" db:"facility"`
	FacilityAddress               string  `json:"facility_address" db:"facility_address"`
	FacilityCity                  string  `json:"facility_city" db:"facility_city"`
	FacilityCode                  string  `json:"facility_code" db:"facility_code"`
	FacilityMunicipality          string  `json:"facility_municipality" db:"facility_municipality"`
	FacilityPostalZipcode         string  `json:"facility_postal_zipcode" db:"facility_postal_zipcode"`
	FoodServiceDescription        string  `json:"food_service_description" db:"food_service_description"`
	FoodServiceType               string  `json:"food_service_type" db:"food_service_type"`
	FsFacilityState               string  `json:"fs_facility_state" db:"fs_facility_state"`
	InspectionComments            string  `json:"inspection_comments" db:"inspection_comments"`
	InspectionType                string  `json:"inspection_type" db:"inspection_type"`
	InspectorID                   string  `json:"inspector_id" db:"inspector_id"`
	Latitude                      float64 `json:"latitude" db:"latitude"`
	LocalHealthDepartment         string  `json:"local_health_department" db:"local_health_department"`
	Longitude                     float64 `json:"longitude" db:"longitude"`
	NysdohGazetteer1980           int     `json:"nysdoh_gazetteer_1980" db:"nysdoh_gazetteer_1980"`
	NysHealthOperationID          string  `json:"nys_health_operation_id" db:"nys_health_operation_id"`
	OperationName                 string  `json:"operation_name" db:"operation_name"`
	PermitExpirationDate          time.Time    `json:"permit_expiration_date" db:"permit_expiration_date"`
	PermittedCorpName             string  `json:"permitted_corp_name" db:"permitted_corp_name"`
	PermittedDba                  string  `json:"permitted_dba" db:"permitted_dba"`
	PermOperatorFirstName         string  `json:"perm_operator_first_name" db:"perm_operator_first_name"`
	PermOperatorLastName          string  `json:"perm_operator_last_name" db:"perm_operator_last_name"`
	TotalNumCriticalViolations    int     `json:"total_num_critical_violations" db:"total_num_critical_violations"`
	TotalNumCritNotCorrected      int     `json:"total_num_crit_not_corrected" db:"total_num_crit_not_corrected"`
	TotalNumNoncriticalViolations int     `json:"total_num_noncritical_violations" db:"total_num_noncritical_violations"`
	ViolationDescription          string  `json:"violation_description" db:"violation_description"`
	ViolationItem                 string  `json:"violation_item" db:"violation_item"`
}

func (r Restaurant) string() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

type Restaurants []Restaurant

func (r Restaurants) string() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}
