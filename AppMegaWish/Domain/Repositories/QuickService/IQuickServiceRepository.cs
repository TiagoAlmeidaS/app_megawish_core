using Domain.Entities.QuickService;
using Domain.Shared.GenericRepositories;

namespace Domain.Repositories.QuickService;

public interface IQuickServiceRepository: IGetRepository<QuickServiceEntity>,
    IInsertRepository<QuickServiceEntity>,
    IUpdateRepository<QuickServiceEntity>,
    IGetByIdRepository<QuickServiceEntity>,
    IDeleteByIdRepository<QuickServiceEntity>
{
    
}