using Domain.Entities.HistoricService;
using Domain.Shared.GenericRepositories;

namespace Domain.Repositories.HistoricService;

public interface IHistoricServiceRepository: IGetRepository<HistoricServiceEntity>,
    IInsertRepository<HistoricServiceEntity>,
    IUpdateRepository<HistoricServiceEntity>,
    IGetByIdRepository<HistoricServiceEntity>,
    IDeleteByIdRepository<HistoricServiceEntity>
{
    
}