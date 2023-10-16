using Domain.Shared;

namespace Domain.Entities.Service;

public class ServiceEntity : AggregateRoot
{
    public int ID { get; set; }
    public string Name { get; set; }
    public string Description { get; set; }
    public int ServiceTypeId { get; set; }
}