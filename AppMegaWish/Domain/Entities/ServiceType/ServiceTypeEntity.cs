using Domain.Shared;

namespace Domain.Entities.ServiceType;

public class ServiceTypeEntity : AggregateRoot
{
    public ServiceTypeEntity(int id, string name, string description)
    {
        ID = id;
        Name = name;
        Description = description;
    }

    public int ID { get; set; }
    public string Name { get; set; }
    public string Description { get; set; }
}