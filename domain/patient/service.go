package patient

import (
	"github.com/jackc/pgx/v4/pgxpool"
	pkgCrypto "zmed_patient_manager/pkg/crypto"
)

type service struct {
	dbConn *pgxpool.Pool
	crypto *pkgCrypto.GcmCrypto
}

func New(dbConn *pgxpool.Pool, crypto *pkgCrypto.GcmCrypto) *service {
	return &service{
		dbConn: dbConn,
		crypto: crypto,
	}
}
