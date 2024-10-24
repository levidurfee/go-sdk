package infisical

import (
	api "github.com/infisical/go-sdk/packages/api/dynamic_secrets"
	"github.com/infisical/go-sdk/packages/models"
)

type ListDynamicSecretLeasesOptions = api.ListDynamicSecretLeaseV1Request
type CreateDynamicSecretLeaseOptions = api.CreateDynamicSecretLeaseV1Request
type DeleteDynamicSecretLeaseOptions = api.DeleteDynamicSecretLeaseV1Request
type GetDynamicSecretLeaseByIdOptions = api.GetDynamicSecretLeaseByIdV1Request
type RenewDynamicSecretLeaseOptions = api.RenewDynamicSecretLeaseV1Request
type ListDynamicSecretsRootCredentialsOptions = api.ListDynamicSecretsV1Request
type GetDynamicSecretRootCredentialByNameOptions = api.GetDynamicSecretByNameV1Request

type DynamicSecretsInterface interface {
	ListRootCredentials(options ListDynamicSecretsRootCredentialsOptions) ([]models.DynamicSecret, error)
	GetRootCredentialByName(options GetDynamicSecretRootCredentialByNameOptions) (models.DynamicSecret, error)
	ListLeases(options ListDynamicSecretLeasesOptions) ([]models.DynamicSecretLease, error)
	CreateLease(options CreateDynamicSecretLeaseOptions) (map[string]any, models.DynamicSecret, models.DynamicSecretLease, error)
	GetLeaseById(options GetDynamicSecretLeaseByIdOptions) (models.DynamicSecretLeaseWithDynamicSecret, error)
	DeleteLeaseById(options DeleteDynamicSecretLeaseOptions) (models.DynamicSecretLease, error)
	RenewLeaseById(options RenewDynamicSecretLeaseOptions) (models.DynamicSecretLease, error)
}

type DynamicSecrets struct {
	client *InfisicalClient
}

func (f *DynamicSecrets) ListRootCredentials(options ListDynamicSecretsRootCredentialsOptions) ([]models.DynamicSecret, error) {
	res, err := api.CallListDynamicSecretsV1(f.client.httpClient, options)

	if err != nil {
		return nil, err
	}

	return res.DynamicSecrets, nil
}

func (f *DynamicSecrets) GetRootCredentialByName(options GetDynamicSecretRootCredentialByNameOptions) (models.DynamicSecret, error) {
	res, err := api.CallGetDynamicSecretByNameV1(f.client.httpClient, options)

	if err != nil {
		return models.DynamicSecret{}, err
	}

	return res.DynamicSecret, nil
}

func (f *DynamicSecrets) ListLeases(options ListDynamicSecretLeasesOptions) ([]models.DynamicSecretLease, error) {
	res, err := api.CallListDynamicSecretLeaseV1(f.client.httpClient, options)

	if err != nil {
		return nil, err
	}

	return res.Leases, nil
}

func (f *DynamicSecrets) CreateLease(options CreateDynamicSecretLeaseOptions) (map[string]any, models.DynamicSecret, models.DynamicSecretLease, error) {
	res, err := api.CallCreateDynamicSecretLeaseV1(f.client.httpClient, options)

	if err != nil {
		return nil, models.DynamicSecret{}, models.DynamicSecretLease{}, err
	}

	return res.Data, res.DynamicSecret, res.Lease, nil
}

func (f *DynamicSecrets) GetLeaseById(options GetDynamicSecretLeaseByIdOptions) (models.DynamicSecretLeaseWithDynamicSecret, error) {
	res, err := api.CallGetByDynamicSecretByIdLeaseV1(f.client.httpClient, options)

	if err != nil {
		return models.DynamicSecretLeaseWithDynamicSecret{}, err
	}

	return res.Lease, nil
}

func (f *DynamicSecrets) DeleteLeaseById(options DeleteDynamicSecretLeaseOptions) (models.DynamicSecretLease, error) {
	res, err := api.CallDeleteDynamicSecretLeaseV1(f.client.httpClient, options)

	if err != nil {
		return models.DynamicSecretLease{}, err
	}

	return res.Lease, nil
}

func (f *DynamicSecrets) RenewLeaseById(options RenewDynamicSecretLeaseOptions) (models.DynamicSecretLease, error) {
	res, err := api.CallRenewDynamicSecretLeaseV1(f.client.httpClient, options)

	if err != nil {
		return models.DynamicSecretLease{}, err
	}

	return res.Lease, nil
}

func NewDynamicSecrets(client *InfisicalClient) DynamicSecretsInterface {
	return &DynamicSecrets{client: client}
}
