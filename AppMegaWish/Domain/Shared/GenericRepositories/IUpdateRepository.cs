namespace Domain.Shared.GenericRepositories;

public interface IUpdateRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<string> Update(TAggregate aggregateRoot, CancellationToken cancellationToken);
}