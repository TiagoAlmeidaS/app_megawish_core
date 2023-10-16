namespace Domain.Shared.GenericRepositories;

public interface IGetByIdRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<TAggregate> GetById(int id, CancellationToken cancellationToken);
}