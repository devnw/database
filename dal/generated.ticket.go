package dal

//**********************************************************
// GENERATED CODE - DO NOT CHANGE
// This file is generated using scaffolding. Any changes to
// this file will be overwritten on the next build
//**********************************************************

import (
	"encoding/json"
	"time"
)

//**********************************************************
// Struct Declaration
//**********************************************************

// Ticket defines the struct that implements the Ticket interface
type Ticket struct {
	AlertDatevar          *time.Time
	AssignedTovar         *string
	AssignmentGroupvar    *string
	CERFvar               string
	CERFExpirationDatevar time.Time
	CVEReferencesvar      *string
	CVSSvar               *float32
	CloudIDvar            string
	Configsvar            string
	CreatedDatevar        *time.Time
	DBCreatedDatevar      time.Time
	DBUpdatedDatevar      *time.Time
	Descriptionvar        *string
	DeviceIDvar           string
	DueDatevar            *time.Time
	GroupIDvar            string
	HostNamevar           *string
	IDvar                 int
	IPAddressvar          *string
	Labelsvar             *string
	LastCheckedvar        *time.Time
	MacAddressvar         *string
	MethodOfDiscoveryvar  *string
	OSDetailedvar         *string
	OperatingSystemvar    *string
	OrgCodevar            *string
	OrganizationIDvar     string
	Priorityvar           *string
	Projectvar            *string
	ReportedByvar         *string
	ResolutionDatevar     *time.Time
	ResolutionStatusvar   *string
	ScanIDvar             int
	ServicePortsvar       *string
	Solutionvar           *string
	Statusvar             *string
	Summaryvar            *string
	TicketTypevar         *string
	Titlevar              string
	UpdatedDatevar        *time.Time
	VendorReferencesvar   *string
	VulnerabilityIDvar    string
	VulnerabilityTitlevar *string
}

//**********************************************************
// Struct Methods
//**********************************************************

// MarshalJSON marshals the struct by converting it to a map
func (myTicket Ticket) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"AlertDate":          myTicket.AlertDatevar,
		"AssignedTo":         myTicket.AssignedTovar,
		"AssignmentGroup":    myTicket.AssignmentGroupvar,
		"CERF":               myTicket.CERFvar,
		"CERFExpirationDate": myTicket.CERFExpirationDatevar,
		"CVEReferences":      myTicket.CVEReferencesvar,
		"CVSS":               myTicket.CVSSvar,
		"CloudID":            myTicket.CloudIDvar,
		"Configs":            myTicket.Configsvar,
		"CreatedDate":        myTicket.CreatedDatevar,
		"DBCreatedDate":      myTicket.DBCreatedDatevar,
		"DBUpdatedDate":      myTicket.DBUpdatedDatevar,
		"Description":        myTicket.Descriptionvar,
		"DeviceID":           myTicket.DeviceIDvar,
		"DueDate":            myTicket.DueDatevar,
		"GroupID":            myTicket.GroupIDvar,
		"HostName":           myTicket.HostNamevar,
		"ID":                 myTicket.IDvar,
		"IPAddress":          myTicket.IPAddressvar,
		"Labels":             myTicket.Labelsvar,
		"LastChecked":        myTicket.LastCheckedvar,
		"MacAddress":         myTicket.MacAddressvar,
		"MethodOfDiscovery":  myTicket.MethodOfDiscoveryvar,
		"OSDetailed":         myTicket.OSDetailedvar,
		"OperatingSystem":    myTicket.OperatingSystemvar,
		"OrgCode":            myTicket.OrgCodevar,
		"OrganizationID":     myTicket.OrganizationIDvar,
		"Priority":           myTicket.Priorityvar,
		"Project":            myTicket.Projectvar,
		"ReportedBy":         myTicket.ReportedByvar,
		"ResolutionDate":     myTicket.ResolutionDatevar,
		"ResolutionStatus":   myTicket.ResolutionStatusvar,
		"ScanID":             myTicket.ScanIDvar,
		"ServicePorts":       myTicket.ServicePortsvar,
		"Solution":           myTicket.Solutionvar,
		"Status":             myTicket.Statusvar,
		"Summary":            myTicket.Summaryvar,
		"TicketType":         myTicket.TicketTypevar,
		"Title":              myTicket.Titlevar,
		"UpdatedDate":        myTicket.UpdatedDatevar,
		"VendorReferences":   myTicket.VendorReferencesvar,
		"VulnerabilityID":    myTicket.VulnerabilityIDvar,
		"VulnerabilityTitle": myTicket.VulnerabilityTitlevar,
	})
}

// AlertDate returns the AlertDate parameter from the Ticket struct
func (myTicket *Ticket) AlertDate() (param *time.Time) {
	return myTicket.AlertDatevar
}

// AssignedTo returns the AssignedTo parameter from the Ticket struct
func (myTicket *Ticket) AssignedTo() (param *string) {
	return myTicket.AssignedTovar
}

// AssignmentGroup returns the AssignmentGroup parameter from the Ticket struct
func (myTicket *Ticket) AssignmentGroup() (param *string) {
	return myTicket.AssignmentGroupvar
}

// CERF returns the CERF parameter from the Ticket struct
func (myTicket *Ticket) CERF() (param string) {
	return myTicket.CERFvar
}

// CERFExpirationDate returns the CERFExpirationDate parameter from the Ticket struct
func (myTicket *Ticket) CERFExpirationDate() (param time.Time) {
	return myTicket.CERFExpirationDatevar
}

// CVEReferences returns the CVEReferences parameter from the Ticket struct
func (myTicket *Ticket) CVEReferences() (param *string) {
	return myTicket.CVEReferencesvar
}

// CVSS returns the CVSS parameter from the Ticket struct
func (myTicket *Ticket) CVSS() (param *float32) {
	return myTicket.CVSSvar
}

// CloudID returns the CloudID parameter from the Ticket struct
func (myTicket *Ticket) CloudID() (param string) {
	return myTicket.CloudIDvar
}

// Configs returns the Configs parameter from the Ticket struct
func (myTicket *Ticket) Configs() (param string) {
	return myTicket.Configsvar
}

// CreatedDate returns the CreatedDate parameter from the Ticket struct
func (myTicket *Ticket) CreatedDate() (param *time.Time) {
	return myTicket.CreatedDatevar
}

// DBCreatedDate returns the DBCreatedDate parameter from the Ticket struct
func (myTicket *Ticket) DBCreatedDate() (param time.Time) {
	return myTicket.DBCreatedDatevar
}

// DBUpdatedDate returns the DBUpdatedDate parameter from the Ticket struct
func (myTicket *Ticket) DBUpdatedDate() (param *time.Time) {
	return myTicket.DBUpdatedDatevar
}

// Description returns the Description parameter from the Ticket struct
func (myTicket *Ticket) Description() (param *string) {
	return myTicket.Descriptionvar
}

// DeviceID returns the DeviceID parameter from the Ticket struct
func (myTicket *Ticket) DeviceID() (param string) {
	return myTicket.DeviceIDvar
}

// DueDate returns the DueDate parameter from the Ticket struct
func (myTicket *Ticket) DueDate() (param *time.Time) {
	return myTicket.DueDatevar
}

// GroupID returns the GroupID parameter from the Ticket struct
func (myTicket *Ticket) GroupID() (param string) {
	return myTicket.GroupIDvar
}

// HostName returns the HostName parameter from the Ticket struct
func (myTicket *Ticket) HostName() (param *string) {
	return myTicket.HostNamevar
}

// ID returns the ID parameter from the Ticket struct
func (myTicket *Ticket) ID() (param int) {
	return myTicket.IDvar
}

// IPAddress returns the IPAddress parameter from the Ticket struct
func (myTicket *Ticket) IPAddress() (param *string) {
	return myTicket.IPAddressvar
}

// Labels returns the Labels parameter from the Ticket struct
func (myTicket *Ticket) Labels() (param *string) {
	return myTicket.Labelsvar
}

// LastChecked returns the LastChecked parameter from the Ticket struct
func (myTicket *Ticket) LastChecked() (param *time.Time) {
	return myTicket.LastCheckedvar
}

// MacAddress returns the MacAddress parameter from the Ticket struct
func (myTicket *Ticket) MacAddress() (param *string) {
	return myTicket.MacAddressvar
}

// MethodOfDiscovery returns the MethodOfDiscovery parameter from the Ticket struct
func (myTicket *Ticket) MethodOfDiscovery() (param *string) {
	return myTicket.MethodOfDiscoveryvar
}

// OSDetailed returns the OSDetailed parameter from the Ticket struct
func (myTicket *Ticket) OSDetailed() (param *string) {
	return myTicket.OSDetailedvar
}

// OperatingSystem returns the OperatingSystem parameter from the Ticket struct
func (myTicket *Ticket) OperatingSystem() (param *string) {
	return myTicket.OperatingSystemvar
}

// OrgCode returns the OrgCode parameter from the Ticket struct
func (myTicket *Ticket) OrgCode() (param *string) {
	return myTicket.OrgCodevar
}

// OrganizationID returns the OrganizationID parameter from the Ticket struct
func (myTicket *Ticket) OrganizationID() (param string) {
	return myTicket.OrganizationIDvar
}

// Priority returns the Priority parameter from the Ticket struct
func (myTicket *Ticket) Priority() (param *string) {
	return myTicket.Priorityvar
}

// Project returns the Project parameter from the Ticket struct
func (myTicket *Ticket) Project() (param *string) {
	return myTicket.Projectvar
}

// ReportedBy returns the ReportedBy parameter from the Ticket struct
func (myTicket *Ticket) ReportedBy() (param *string) {
	return myTicket.ReportedByvar
}

// ResolutionDate returns the ResolutionDate parameter from the Ticket struct
func (myTicket *Ticket) ResolutionDate() (param *time.Time) {
	return myTicket.ResolutionDatevar
}

// ResolutionStatus returns the ResolutionStatus parameter from the Ticket struct
func (myTicket *Ticket) ResolutionStatus() (param *string) {
	return myTicket.ResolutionStatusvar
}

// ScanID returns the ScanID parameter from the Ticket struct
func (myTicket *Ticket) ScanID() (param int) {
	return myTicket.ScanIDvar
}

// ServicePorts returns the ServicePorts parameter from the Ticket struct
func (myTicket *Ticket) ServicePorts() (param *string) {
	return myTicket.ServicePortsvar
}

// Solution returns the Solution parameter from the Ticket struct
func (myTicket *Ticket) Solution() (param *string) {
	return myTicket.Solutionvar
}

// Status returns the Status parameter from the Ticket struct
func (myTicket *Ticket) Status() (param *string) {
	return myTicket.Statusvar
}

// Summary returns the Summary parameter from the Ticket struct
func (myTicket *Ticket) Summary() (param *string) {
	return myTicket.Summaryvar
}

// TicketType returns the TicketType parameter from the Ticket struct
func (myTicket *Ticket) TicketType() (param *string) {
	return myTicket.TicketTypevar
}

// Title returns the Title parameter from the Ticket struct
func (myTicket *Ticket) Title() (param string) {
	return myTicket.Titlevar
}

// UpdatedDate returns the UpdatedDate parameter from the Ticket struct
func (myTicket *Ticket) UpdatedDate() (param *time.Time) {
	return myTicket.UpdatedDatevar
}

// VendorReferences returns the VendorReferences parameter from the Ticket struct
func (myTicket *Ticket) VendorReferences() (param *string) {
	return myTicket.VendorReferencesvar
}

// VulnerabilityID returns the VulnerabilityID parameter from the Ticket struct
func (myTicket *Ticket) VulnerabilityID() (param string) {
	return myTicket.VulnerabilityIDvar
}

// VulnerabilityTitle returns the VulnerabilityTitle parameter from the Ticket struct
func (myTicket *Ticket) VulnerabilityTitle() (param *string) {
	return myTicket.VulnerabilityTitlevar
}

// SetAssignedTo sets the AssignedTo parameter from the Ticket struct
func (myTicket *Ticket) SetAssignedTo(val string) {
	myTicket.AssignedTovar = &val
}

// SetAssignmentGroup sets the AssignmentGroup parameter from the Ticket struct
func (myTicket *Ticket) SetAssignmentGroup(val string) {
	myTicket.AssignmentGroupvar = &val
}

// SetCERF sets the CERF parameter from the Ticket struct
func (myTicket *Ticket) SetCERF(val string) {
	myTicket.CERFvar = val
}

// SetCVEReferences sets the CVEReferences parameter from the Ticket struct
func (myTicket *Ticket) SetCVEReferences(val string) {
	myTicket.CVEReferencesvar = &val
}

// SetCloudID sets the CloudID parameter from the Ticket struct
func (myTicket *Ticket) SetCloudID(val string) {
	myTicket.CloudIDvar = val
}

// SetConfigs sets the Configs parameter from the Ticket struct
func (myTicket *Ticket) SetConfigs(val string) {
	myTicket.Configsvar = val
}

// SetDeviceID sets the DeviceID parameter from the Ticket struct
func (myTicket *Ticket) SetDeviceID(val string) {
	myTicket.DeviceIDvar = val
}

// SetGroupID sets the GroupID parameter from the Ticket struct
func (myTicket *Ticket) SetGroupID(val string) {
	myTicket.GroupIDvar = val
}

// SetHostName sets the HostName parameter from the Ticket struct
func (myTicket *Ticket) SetHostName(val string) {
	myTicket.HostNamevar = &val
}

// SetIPAddress sets the IPAddress parameter from the Ticket struct
func (myTicket *Ticket) SetIPAddress(val string) {
	myTicket.IPAddressvar = &val
}

// SetLabels sets the Labels parameter from the Ticket struct
func (myTicket *Ticket) SetLabels(val string) {
	myTicket.Labelsvar = &val
}

// SetMacAddress sets the MacAddress parameter from the Ticket struct
func (myTicket *Ticket) SetMacAddress(val string) {
	myTicket.MacAddressvar = &val
}

// SetMethodOfDiscovery sets the MethodOfDiscovery parameter from the Ticket struct
func (myTicket *Ticket) SetMethodOfDiscovery(val string) {
	myTicket.MethodOfDiscoveryvar = &val
}

// SetOSDetailed sets the OSDetailed parameter from the Ticket struct
func (myTicket *Ticket) SetOSDetailed(val string) {
	myTicket.OSDetailedvar = &val
}

// SetOperatingSystem sets the OperatingSystem parameter from the Ticket struct
func (myTicket *Ticket) SetOperatingSystem(val string) {
	myTicket.OperatingSystemvar = &val
}

// SetOrgCode sets the OrgCode parameter from the Ticket struct
func (myTicket *Ticket) SetOrgCode(val string) {
	myTicket.OrgCodevar = &val
}

// SetOrganizationID sets the OrganizationID parameter from the Ticket struct
func (myTicket *Ticket) SetOrganizationID(val string) {
	myTicket.OrganizationIDvar = val
}

// SetPriority sets the Priority parameter from the Ticket struct
func (myTicket *Ticket) SetPriority(val string) {
	myTicket.Priorityvar = &val
}

// SetProject sets the Project parameter from the Ticket struct
func (myTicket *Ticket) SetProject(val string) {
	myTicket.Projectvar = &val
}

// SetReportedBy sets the ReportedBy parameter from the Ticket struct
func (myTicket *Ticket) SetReportedBy(val string) {
	myTicket.ReportedByvar = &val
}

// SetResolutionStatus sets the ResolutionStatus parameter from the Ticket struct
func (myTicket *Ticket) SetResolutionStatus(val string) {
	myTicket.ResolutionStatusvar = &val
}

// SetServicePorts sets the ServicePorts parameter from the Ticket struct
func (myTicket *Ticket) SetServicePorts(val string) {
	myTicket.ServicePortsvar = &val
}

// SetStatus sets the Status parameter from the Ticket struct
func (myTicket *Ticket) SetStatus(val string) {
	myTicket.Statusvar = &val
}

// SetSummary sets the Summary parameter from the Ticket struct
func (myTicket *Ticket) SetSummary(val string) {
	myTicket.Summaryvar = &val
}

// SetTicketType sets the TicketType parameter from the Ticket struct
func (myTicket *Ticket) SetTicketType(val string) {
	myTicket.TicketTypevar = &val
}

// SetTitle sets the Title parameter from the Ticket struct
func (myTicket *Ticket) SetTitle(val string) {
	myTicket.Titlevar = val
}

// SetVulnerabilityID sets the VulnerabilityID parameter from the Ticket struct
func (myTicket *Ticket) SetVulnerabilityID(val string) {
	myTicket.VulnerabilityIDvar = val
}

// SetVulnerabilityTitle sets the VulnerabilityTitle parameter from the Ticket struct
func (myTicket *Ticket) SetVulnerabilityTitle(val string) {
	myTicket.VulnerabilityTitlevar = &val
}
