namespace Domain.Shared;

public class Entity
{
    public Entity()
    {
        Guid = Guid.NewGuid();
        UpdatedAt = DateTime.UtcNow;
        CreatedAt = DateTime.UtcNow;
    }

    public Guid Guid { get; set; }
    public DateTime UpdatedAt { get; set; }
    public DateTime CreatedAt { get; set; }
}