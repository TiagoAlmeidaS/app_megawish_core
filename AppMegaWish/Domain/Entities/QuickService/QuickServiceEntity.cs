using Domain.Shared;

namespace Domain.Entities.QuickService;

public class QuickServiceEntity : AggregateRoot
{
    public QuickServiceEntity(int id, int serviceTypeId)
    {
        ID = id;
        ServiceTypeID = serviceTypeId;
    }

    public int ID { get; set; }
    public int ServiceTypeID { get; set; }
}