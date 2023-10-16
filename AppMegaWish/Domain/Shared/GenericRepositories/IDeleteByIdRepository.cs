namespace Domain.Shared.GenericRepositories;

public interface IDeleteByIdRepository<TAggregate> : IRepository
    where TAggregate : AggregateRoot
{
    public Task<TAggregate> DeleteById(int id, CancellationToken cancellationToken);
}