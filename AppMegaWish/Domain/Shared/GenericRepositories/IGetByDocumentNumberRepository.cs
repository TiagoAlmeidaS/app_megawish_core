namespace Domain.Shared.GenericRepositories;

public interface IGetByDocumentNumberRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<TAggregate> GetByDocumentNumber(string documentNumber, CancellationToken cancellationToken);
}