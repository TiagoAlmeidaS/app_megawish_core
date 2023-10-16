namespace Domain.Shared.GenericRepositories;

public interface IInsertRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<string> Insert(TAggregate aggregate, CancellationToken cancellationToken);
}