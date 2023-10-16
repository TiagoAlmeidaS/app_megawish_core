using Domain.Shared;

namespace Domain.Entities.HistoricService;

public class HistoricServiceEntity : AggregateRoot
{
    public HistoricServiceEntity(int id, DateTime dataInicio, DateTime dataFim, float valor, int serviceId,
        int statusId, int clientId)
    {
        ID = id;
        DataInicio = dataInicio;
        DataFim = dataFim;
        Valor = valor;
        ServiceID = serviceId;
        StatusID = statusId;
        ClientID = clientId;
    }

    public int ID { get; set; }
    public DateTime DataInicio { get; set; }
    public DateTime DataFim { get; set; }
    public float Valor { get; set; }
    public int ServiceID { get; set; }
    public int StatusID { get; set; }
    public int ClientID { get; set; }
}