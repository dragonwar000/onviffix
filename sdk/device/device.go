package device

//go:generate go run github.com/dragonwar000/onviffix/sdk/codegen device device GetServices
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetServiceCapabilities
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDeviceInformation
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetSystemDateAndTime
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetSystemDateAndTime
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetSystemFactoryDefault
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device UpgradeSystemFirmware
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SystemReboot
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device RestoreSystem
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetSystemBackup
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetSystemLog
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetSystemSupportInformation
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetScopes
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetScopes
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device AddScopes
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device RemoveScopes
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDiscoveryMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetDiscoveryMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetRemoteDiscoveryMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetRemoteDiscoveryMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDPAddresses
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetEndpointReference
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetRemoteUser
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetRemoteUser
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetUsers
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device CreateUsers
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device DeleteUsers
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetUser
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetWsdlUrl
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetCapabilities
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetHostname
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetHostname
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetHostnameFromDHCP
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDNS
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetDNS
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetNTP
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetNTP
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDynamicDNS
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetDynamicDNS
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetNetworkInterfaces
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetNetworkInterfaces
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetNetworkProtocols
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetNetworkProtocols
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetNetworkDefaultGateway
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetNetworkDefaultGateway
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetZeroConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetZeroConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetIPAddressFilter
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetIPAddressFilter
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device AddIPAddressFilter
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device RemoveIPAddressFilter
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetAccessPolicy
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetAccessPolicy
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device CreateCertificate
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetCertificates
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetCertificatesStatus
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetCertificatesStatus
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device DeleteCertificates
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetPkcs10Request
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device LoadCertificates
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetClientCertificateMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetClientCertificateMode
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetRelayOutputs
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetRelayOutputSettings
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetRelayOutputState
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SendAuxiliaryCommand
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetCACertificates
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device LoadCertificateWithPrivateKey
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetCertificateInformation
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device LoadCACertificates
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device CreateDot1XConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetDot1XConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDot1XConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDot1XConfigurations
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device DeleteDot1XConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDot11Capabilities
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetDot11Status
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device ScanAvailableDot11Networks
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetSystemUris
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device StartFirmwareUpgrade
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device StartSystemRestore
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetStorageConfigurations
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device CreateStorageConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetStorageConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetStorageConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device DeleteStorageConfiguration
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device GetGeoLocation
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device SetGeoLocation
//go:generate go run github.com/dragonwar000/onviffix/onviffix/sdk/codegen device device DeleteGeoLocation
