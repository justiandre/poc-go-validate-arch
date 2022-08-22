package main

import (
	"log"

	domain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
	restdomain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01/entrypoint/rest"
	httpdomain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01/repository/http"
	sqldomain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01/repository/sql"
	servicedomain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01/service"
	usecasedomain01 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01/usecase"
	domain02 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
	sqldomain02 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02/repository/sql"
	servicedomain02 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02/service"
	domain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03"
	restdomain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03/entrypoint/rest"
	sqldomain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03/repository/sql"
	servicedomain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03/service"
	usecasedomain03 "github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03/usecase"
)

// TODO https://github.com/fdaines/arch-go

func main() {
	log.Println("this is just a poc, func: main.go#main - orchestrates all modules, test only!!!")

	//domain01
	entityDomain01 := domain01.Domain01{ID: "01", Value01: "Domain01_01", Value02: "Domain01_02"}
	restDomain01 := createRestDomain01()
	restDomain01.FindByID("01")
	restDomain01.Save(entityDomain01)
	restDomain01.DeleteByID("01")

	//domain02
	entityDomain02 := domain02.Domain02{ID: "02", Value01: "Domain02_01", Value02: "Domain02_02"}
	serviceDomain02 := createServiceDomain02()

	//domain03
	entityDomain03 := domain03.Domain03{ID: "03", Value01: "Domain03_01", Value02: "Domain03_02"}
	restDomain03 := createRestDomain03(serviceDomain02)
	restDomain03.Save(entityDomain03, entityDomain02)
}

func createRestDomain01() restdomain01.Rest {
	httpDomain01 := httpdomain01.New()
	sqlDomain01 := sqldomain01.New()
	type repositoryInlineDomain01 struct {
		httpdomain01.Http
		sqldomain01.Sql
	}
	serviceDomain01 := servicedomain01.New(repositoryInlineDomain01{httpDomain01, sqlDomain01})
	useCaseFindUcDomain01 := usecasedomain01.NewFindUC(serviceDomain01)
	useCaseSaveUcDomain01 := usecasedomain01.NewSaveUC(serviceDomain01)
	useCaseDeleteUcDomain01 := usecasedomain01.NewDeleteUC(serviceDomain01)
	type useCaseInlineDomain01 struct {
		usecasedomain01.FindUC
		usecasedomain01.SaveUC
		usecasedomain01.DeleteUC
	}
	return restdomain01.New(useCaseInlineDomain01{useCaseFindUcDomain01, useCaseSaveUcDomain01, useCaseDeleteUcDomain01})
}

func createServiceDomain02() servicedomain02.Service {
	sqlDomain02 := sqldomain02.New()
	return servicedomain02.New(sqlDomain02)
}

func createRestDomain03(serviceDomain02 servicedomain02.Service) restdomain03.Rest {
	sqlDomain03 := sqldomain03.New()
	serviceDomain03 := servicedomain03.New(sqlDomain03)
	useCaseDomain03 := usecasedomain03.NewSaveUC(serviceDomain03, serviceDomain02)
	return restdomain03.New(useCaseDomain03)
}
