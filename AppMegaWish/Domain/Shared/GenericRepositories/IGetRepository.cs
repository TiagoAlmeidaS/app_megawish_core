namespace Domain.Shared.GenericRepositories;

public interface IGetRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<TAggregate> Get(string guid, CancellationToken cancellationToken);
}