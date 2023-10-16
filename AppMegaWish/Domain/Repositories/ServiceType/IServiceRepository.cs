using Domain.Entities.Service;
using Domain.Entities.ServiceType;
using Domain.Shared.GenericRepositories;

namespace Domain.Repositories.ServiceType;

public interface IServiceTypeRepository : IGetRepository<ServiceTypeEntity>,
    IInsertRepository<ServiceTypeEntity>,
    IUpdateRepository<ServiceTypeEntity>,
    IGetByIdRepository<ServiceTypeEntity>,
    IDeleteByIdRepository<ServiceTypeEntity>
{
}