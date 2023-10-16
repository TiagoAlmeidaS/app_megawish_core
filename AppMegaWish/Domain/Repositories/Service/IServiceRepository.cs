using Domain.Entities.Service;
using Domain.Shared.GenericRepositories;

namespace Domain.Repositories.Service;

public interface IServiceRepository : IGetRepository<ServiceEntity>,
    IInsertRepository<ServiceEntity>,
    IUpdateRepository<ServiceEntity>,
    IGetByIdRepository<ServiceEntity>,
    IDeleteByIdRepository<ServiceEntity>
{
}