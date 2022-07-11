// Package countries supports subdivisions as per ISO 3166-2.
//
// Data has been sourced from <https://www.ip2location.com/free/iso3166-2>. See
// license for further information.
package countries

// TypeSubdivisionTypeCode for Typer interface.
const TypeSubdivisionTypeCode string = "countries.SubdivisionTypeCode"

// TypeSubdivisionType for Typer interface.
const TypeSubdivisionType string = "countries.SubdivisionType"

// Subdivision types
const (
	// SubdivisionTypeUnknown                                                         SubdivisionType = "Unknown"
	SubdivisionTypeUnknown = "unknown"
	// SubdivisionTypeAdministration                                                  SubdivisionTypeCode = "Administration"
	SubdivisionTypeAdministration SubdivisionTypeCode = "Administration"
	// SubdivisionTypeAdministrativeAtoll                                             SubdivisionTypeCode = "Administrative Atoll"
	SubdivisionTypeAdministrativeAtoll SubdivisionTypeCode = "Administrative Atoll"
	// SubdivisionTypeAdministrativeRegion                                            SubdivisionTypeCode = "Administrative Region"
	SubdivisionTypeAdministrativeRegion SubdivisionTypeCode = "Administrative Region"
	// SubdivisionTypeAdministrativeTerritory                                         SubdivisionTypeCode = "Administrative Territory"
	SubdivisionTypeAdministrativeTerritory SubdivisionTypeCode = "Administrative Territory"
	// SubdivisionTypeArcticRegion                                                    SubdivisionTypeCode = "Arctic Region"
	SubdivisionTypeArcticRegion SubdivisionTypeCode = "Arctic Region"
	// SubdivisionTypeArea                                                            SubdivisionTypeCode = "Area"
	SubdivisionTypeArea SubdivisionTypeCode = "Area"
	// SubdivisionTypeAutonomousCity                                                  SubdivisionTypeCode = "Autonomous City"
	SubdivisionTypeAutonomousCity SubdivisionTypeCode = "Autonomous City"
	// SubdivisionTypeAutonomousCommunity                                             SubdivisionTypeCode = "Autonomous Community"
	SubdivisionTypeAutonomousCommunity SubdivisionTypeCode = "Autonomous Community"
	// SubdivisionTypeAutonomousDistrict                                              SubdivisionTypeCode = "Autonomous District"
	SubdivisionTypeAutonomousDistrict SubdivisionTypeCode = "Autonomous District"
	// SubdivisionTypeAutonomousMunicipality                                          SubdivisionTypeCode = "Autonomous Municipality"
	SubdivisionTypeAutonomousMunicipality SubdivisionTypeCode = "Autonomous Municipality"
	// SubdivisionTypeAutonomousProvince                                              SubdivisionTypeCode = "Autonomous Province"
	SubdivisionTypeAutonomousProvince SubdivisionTypeCode = "Autonomous Province"
	// SubdivisionTypeAutonomousRegion                                                SubdivisionTypeCode = "Autonomous Region"
	SubdivisionTypeAutonomousRegion SubdivisionTypeCode = "Autonomous Region"
	// SubdivisionTypeAutonomousRepublic                                              SubdivisionTypeCode = "Autonomous Republic"
	SubdivisionTypeAutonomousRepublic SubdivisionTypeCode = "Autonomous Republic"
	// SubdivisionTypeAutonomousSector                                                SubdivisionTypeCode = "Autonomous Sector"
	SubdivisionTypeAutonomousSector SubdivisionTypeCode = "Autonomous Sector"
	// SubdivisionTypeAutonomousTerritorialUnit                                       SubdivisionTypeCode = "Autonomous Territorial Unit"
	SubdivisionTypeAutonomousTerritorialUnit SubdivisionTypeCode = "Autonomous Territorial Unit"
	// SubdivisionTypeBorough                                                         SubdivisionTypeCode = "Borough"
	SubdivisionTypeBorough SubdivisionTypeCode = "Borough"
	// SubdivisionTypeCanton                                                          SubdivisionTypeCode = "Canton"
	SubdivisionTypeCanton SubdivisionTypeCode = "Canton"
	// SubdivisionTypeCapitalCity                                                     SubdivisionTypeCode = "Capital City"
	SubdivisionTypeCapitalCity SubdivisionTypeCode = "Capital City"
	// SubdivisionTypeCapitalDistrict                                                 SubdivisionTypeCode = "Capital District"
	SubdivisionTypeCapitalDistrict SubdivisionTypeCode = "Capital District"
	// SubdivisionTypeCapitalMetropolitanCity                                         SubdivisionTypeCode = "Capital Metropolitan City"
	SubdivisionTypeCapitalMetropolitanCity SubdivisionTypeCode = "Capital Metropolitan City"
	// SubdivisionTypeCapitalTerritory                                                SubdivisionTypeCode = "Capital Territory"
	SubdivisionTypeCapitalTerritory SubdivisionTypeCode = "Capital Territory"
	// SubdivisionTypeChainsOfIslands                                                 SubdivisionTypeCode = "Chains (of Islands)"
	SubdivisionTypeChainsOfIslands SubdivisionTypeCode = "Chains (of Islands)"
	// SubdivisionTypeCity                                                            SubdivisionTypeCode = "City"
	SubdivisionTypeCity SubdivisionTypeCode = "City"
	// SubdivisionTypeCityCorporation                                                 SubdivisionTypeCode = "City Corporation"
	SubdivisionTypeCityCorporation SubdivisionTypeCode = "City Corporation"
	// SubdivisionTypeCityWithCountyRights                                            SubdivisionTypeCode = "City with County Rights"
	SubdivisionTypeCityWithCountyRights SubdivisionTypeCode = "City with County Rights"
	// SubdivisionTypeCommune                                                         SubdivisionTypeCode = "Commune"
	SubdivisionTypeCommune SubdivisionTypeCode = "Commune"
	// SubdivisionTypeConstitutionalProvince                                          SubdivisionTypeCode = "Constitutional Province"
	SubdivisionTypeConstitutionalProvince SubdivisionTypeCode = "Constitutional Province"
	// SubdivisionTypeCouncilArea                                                     SubdivisionTypeCode = "Council Area"
	SubdivisionTypeCouncilArea SubdivisionTypeCode = "Council Area"
	// SubdivisionTypeCountry                                                         SubdivisionTypeCode = "Country"
	SubdivisionTypeCountry SubdivisionTypeCode = "Country"
	// SubdivisionTypeCounty                                                          SubdivisionTypeCode = "County"
	SubdivisionTypeCounty SubdivisionTypeCode = "County"
	// SubdivisionTypeDepartment                                                      SubdivisionTypeCode = "Department"
	SubdivisionTypeDepartment SubdivisionTypeCode = "Department"
	// SubdivisionTypeDependency                                                      SubdivisionTypeCode = "Dependency"
	SubdivisionTypeDependency SubdivisionTypeCode = "Dependency"
	// SubdivisionTypeDevelopmentRegion                                               SubdivisionTypeCode = "Development Region"
	SubdivisionTypeDevelopmentRegion SubdivisionTypeCode = "Development Region"
	// SubdivisionTypeDistrict                                                        SubdivisionTypeCode = "District"
	SubdivisionTypeDistrict SubdivisionTypeCode = "District"
	// SubdivisionTypeDivision                                                        SubdivisionTypeCode = "Division"
	SubdivisionTypeDivision SubdivisionTypeCode = "Division"
	// SubdivisionTypeEconomicPrefecture                                              SubdivisionTypeCode = "Economic Prefecture"
	SubdivisionTypeEconomicPrefecture SubdivisionTypeCode = "Economic Prefecture"
	// SubdivisionTypeEmirate                                                         SubdivisionTypeCode = "Emirate"
	SubdivisionTypeEmirate SubdivisionTypeCode = "Emirate"
	// SubdivisionTypeEntity                                                          SubdivisionTypeCode = "Entity"
	SubdivisionTypeEntity SubdivisionTypeCode = "Entity"
	// SubdivisionTypeFederalDependency                                               SubdivisionTypeCode = "Federal Dependency"
	SubdivisionTypeFederalDependency SubdivisionTypeCode = "Federal Dependency"
	// SubdivisionTypeFederalDistrict                                                 SubdivisionTypeCode = "Federal District"
	SubdivisionTypeFederalDistrict SubdivisionTypeCode = "Federal District"
	// SubdivisionTypeFederalTerritories                                              SubdivisionTypeCode = "Federal Territories"
	SubdivisionTypeFederalTerritories SubdivisionTypeCode = "Federal Territories"
	// SubdivisionTypeGeographicalEntity                                              SubdivisionTypeCode = "Geographical Entity"
	SubdivisionTypeGeographicalEntity SubdivisionTypeCode = "Geographical Entity"
	// SubdivisionTypeGeographicalRegion                                              SubdivisionTypeCode = "Geographical Region"
	SubdivisionTypeGeographicalRegion SubdivisionTypeCode = "Geographical Region"
	// SubdivisionTypeGeographicalUnit                                                SubdivisionTypeCode = "Geographical Unit"
	SubdivisionTypeGeographicalUnit SubdivisionTypeCode = "Geographical Unit"
	// SubdivisionTypeGovernorate                                                     SubdivisionTypeCode = "Governorate"
	SubdivisionTypeGovernorate SubdivisionTypeCode = "Governorate"
	// SubdivisionTypeIndigenousRegion                                                SubdivisionTypeCode = "Indigenous Region"
	SubdivisionTypeIndigenousRegion SubdivisionTypeCode = "Indigenous Region"
	// SubdivisionTypeIsland                                                          SubdivisionTypeCode = "Island"
	SubdivisionTypeIsland SubdivisionTypeCode = "Island"
	// SubdivisionTypeIslandCouncil                                                   SubdivisionTypeCode = "Island Council"
	SubdivisionTypeIslandCouncil SubdivisionTypeCode = "Island Council"
	// SubdivisionTypeIslandGroup                                                     SubdivisionTypeCode = "Island Group"
	SubdivisionTypeIslandGroup SubdivisionTypeCode = "Island Group"
	// SubdivisionTypeLocalCouncil                                                    SubdivisionTypeCode = "Local Council"
	SubdivisionTypeLocalCouncil SubdivisionTypeCode = "Local Council"
	// SubdivisionTypeLondonBorough                                                   SubdivisionTypeCode = "London Borough"
	SubdivisionTypeLondonBorough SubdivisionTypeCode = "London Borough"
	// SubdivisionTypeMetropolitanCities                                              SubdivisionTypeCode = "Metropolitan Cities"
	SubdivisionTypeMetropolitanCities SubdivisionTypeCode = "Metropolitan Cities"
	// SubdivisionTypeMetropolitanDepartment                                          SubdivisionTypeCode = "Metropolitan Department"
	SubdivisionTypeMetropolitanDepartment SubdivisionTypeCode = "Metropolitan Department"
	// SubdivisionTypeMetropolitanDistrict                                            SubdivisionTypeCode = "Metropolitan District"
	SubdivisionTypeMetropolitanDistrict SubdivisionTypeCode = "Metropolitan District"
	// SubdivisionTypeMetropolitanRegion                                              SubdivisionTypeCode = "Metropolitan Region"
	SubdivisionTypeMetropolitanRegion SubdivisionTypeCode = "Metropolitan Region"
	// SubdivisionTypeMunicipalities                                                  SubdivisionTypeCode = "Municipalities"
	SubdivisionTypeMunicipalities SubdivisionTypeCode = "Municipalities"
	// SubdivisionTypeMunicipality                                                    SubdivisionTypeCode = "Municipality"
	SubdivisionTypeMunicipality SubdivisionTypeCode = "Municipality"
	// SubdivisionTypeNation                                                          SubdivisionTypeCode = "Nation"
	SubdivisionTypeNation SubdivisionTypeCode = "Nation"
	// SubdivisionTypeOblast                                                          SubdivisionTypeCode = "Oblast"
	SubdivisionTypeOblast SubdivisionTypeCode = "Oblast"
	// SubdivisionTypeOutlyingArea                                                    SubdivisionTypeCode = "Outlying Area"
	SubdivisionTypeOutlyingArea SubdivisionTypeCode = "Outlying Area"
	// SubdivisionTypeOverseasDepartment                                              SubdivisionTypeCode = "Overseas Department"
	SubdivisionTypeOverseasDepartment SubdivisionTypeCode = "Overseas Department"
	// SubdivisionTypeOverseasRegion                                                  SubdivisionTypeCode = "Overseas Region"
	SubdivisionTypeOverseasRegion SubdivisionTypeCode = "Overseas Region"
	// SubdivisionTypeOverseasTerritorialCollectivity                                 SubdivisionTypeCode = "Overseas Territorial Collectivity"
	SubdivisionTypeOverseasTerritorialCollectivity SubdivisionTypeCode = "Overseas Territorial Collectivity"
	// SubdivisionTypeParish                                                          SubdivisionTypeCode = "Parish"
	SubdivisionTypeParish SubdivisionTypeCode = "Parish"
	// SubdivisionTypePopularates                                                     SubdivisionTypeCode = "Popularates"
	SubdivisionTypePopularates SubdivisionTypeCode = "Popularates"
	// SubdivisionTypePrefecture                                                      SubdivisionTypeCode = "Prefecture"
	SubdivisionTypePrefecture SubdivisionTypeCode = "Prefecture"
	// SubdivisionTypeProvince                                                        SubdivisionTypeCode = "Province"
	SubdivisionTypeProvince SubdivisionTypeCode = "Province"
	// SubdivisionTypeQuarter                                                         SubdivisionTypeCode = "Quarter"
	SubdivisionTypeQuarter SubdivisionTypeCode = "Quarter"
	// SubdivisionTypeRayon                                                           SubdivisionTypeCode = "Rayon"
	SubdivisionTypeRayon SubdivisionTypeCode = "Rayon"
	// SubdivisionTypeRegion                                                          SubdivisionTypeCode = "Region"
	SubdivisionTypeRegion SubdivisionTypeCode = "Region"
	// SubdivisionTypeRegionalCouncil                                                 SubdivisionTypeCode = "Regional Council"
	SubdivisionTypeRegionalCouncil SubdivisionTypeCode = "Regional Council"
	// SubdivisionTypeRepublic                                                        SubdivisionTypeCode = "Republic"
	SubdivisionTypeRepublic SubdivisionTypeCode = "Republic"
	// SubdivisionTypeRepublicanCity                                                  SubdivisionTypeCode = "Republican City"
	SubdivisionTypeRepublicanCity SubdivisionTypeCode = "Republican City"
	// SubdivisionTypeSelfGovernedPart                                                SubdivisionTypeCode = "Self-governed Part"
	SubdivisionTypeSelfGovernedPart SubdivisionTypeCode = "Self-governed Part"
	// SubdivisionTypeSpecialAdministrativeRegion                                     SubdivisionTypeCode = "Special Administrative Region"
	SubdivisionTypeSpecialAdministrativeRegion SubdivisionTypeCode = "Special Administrative Region"
	// SubdivisionTypeSpecialCity                                                     SubdivisionTypeCode = "Special city"
	SubdivisionTypeSpecialCity SubdivisionTypeCode = "Special City"
	// SubdivisionTypeSpecialDistrict                                                 SubdivisionTypeCode = "Special District"
	SubdivisionTypeSpecialDistrict SubdivisionTypeCode = "Special District"
	// SubdivisionTypeSpecialIslandAuthority                                          SubdivisionTypeCode = "Special Island Authority"
	SubdivisionTypeSpecialIslandAuthority SubdivisionTypeCode = "Special Island Authority"
	// SubdivisionTypeSpecialMunicipality                                             SubdivisionTypeCode = "Special Municipality"
	SubdivisionTypeSpecialMunicipality SubdivisionTypeCode = "Special Municipality"
	// SubdivisionTypeSpecialRegion                                                   SubdivisionTypeCode = "Special Region"
	SubdivisionTypeSpecialRegion SubdivisionTypeCode = "Special Region"
	// SubdivisionTypeSpecialZone                                                     SubdivisionTypeCode = "Special Zone"
	SubdivisionTypeSpecialZone SubdivisionTypeCode = "Special Zone"
	// SubdivisionTypeState                                                           SubdivisionTypeCode = "State"
	SubdivisionTypeState SubdivisionTypeCode = "State"
	// SubdivisionTypeTerritorialUnit                                                 SubdivisionTypeCode = "Territorial Unit"
	SubdivisionTypeTerritorialUnit SubdivisionTypeCode = "Territorial Unit"
	// SubdivisionTypeTerritory                                                       SubdivisionTypeCode = "Territory"
	SubdivisionTypeTerritory SubdivisionTypeCode = "Territory"
	// SubdivisionTypeTownCouncil                                                     SubdivisionTypeCode = "Town Council"
	SubdivisionTypeTownCouncil SubdivisionTypeCode = "Town Council"
	// SubdivisionTypeTwoTierCounty                                                   SubdivisionTypeCode = "Two-tier County"
	SubdivisionTypeTwoTierCounty SubdivisionTypeCode = "Two-tier County"
	// SubdivisionTypeUnionTerritory                                                  SubdivisionTypeCode = "Union Territory"
	SubdivisionTypeUnionTerritory SubdivisionTypeCode = "Union Territory"
	// SubdivisionTypeUnitaryAuthority                                                SubdivisionTypeCode = "Unitary Authority"
	SubdivisionTypeUnitaryAuthority SubdivisionTypeCode = "Unitary Authority"
	// SubdivisionTypeZone                                                            SubdivisionTypeCode = "Zone"
	SubdivisionTypeZone SubdivisionTypeCode = "Zone"
)