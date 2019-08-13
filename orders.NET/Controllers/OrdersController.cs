using System.Collections.Generic;

using Microsoft.AspNetCore.Mvc;
using Halcyon.HAL.Attributes;
using Halcyon.Web.HAL;
using CustomerOrdersApi.Model;
using System.Net;
using Newtonsoft.Json;
using HalKit;
using System.Threading.Tasks;
using System;
using MongoDB.Driver;
using Microsoft.Extensions.Options;
using Microsoft.Extensions.Logging;

using CustomerOrdersApi.Config;
using System.Threading;
using System.IO;

namespace CustomerOrdersApi
{
    [HalModel("http://orders/", true)]
    [HalLink("self", "/orders")]
    [HalLink("profile", "profile/orders")]
    [HalLink("search", "orders/search")]
    public class OrdersModel
    {
        [JsonProperty("page")]
        public ResultPage Page { get; set; } = new ResultPage();
        [JsonIgnore]
        [HalEmbedded("customerOrders")]
        public List<CustomerOrder> Orders { get; set; } = new List<CustomerOrder>();

        public class ResultPage
        {
            [JsonProperty("size")]
            public int Size { get; set; } = 20;
            [JsonProperty("totalElements")]
            public int TotalElements { get; set; }
            [JsonProperty("totalPages")]
            public int TotalPages { get; set; }
            [JsonProperty("number")]
            public int Number { get; set; }
        }
    }

   [Route("/orders")]
   public class ProductsController: Controller
    {
        private readonly AppSettings AppSettings;

        private readonly ILogger<ProductsController> logger;


        private HalClient client = new HalClient(new HalConfiguration());
        private HALAttributeConverter converter = new HALAttributeConverter();

        private MongoClient dbClient;
        private IMongoCollection<CustomerOrder> collection;
        public ProductsController(IOptions<AppSettings> options, ILogger<ProductsController> logger) : base()
        {
            this.AppSettings = options.Value;
            this.logger = logger;
            dbClient = new MongoClient(AppSettings.Data.MongoConnection.ConnectionString);
            IMongoDatabase database = dbClient.GetDatabase(AppSettings.Data.MongoConnection.Database);
            collection = database.GetCollection<CustomerOrder>("customerOrder");
        }

        [HttpGet]
        public IActionResult Get()
        {
            IEnumerator<CustomerOrder> enumerator = collection.AsQueryable<CustomerOrder>().GetEnumerator();
            OrdersModel model = new OrdersModel();

            while (enumerator.MoveNext())
            {
                model.Orders.Add(enumerator.Current);
            }
            model.Page.TotalPages = 1;
            model.Page.TotalElements = model.Orders.Count;
            model.Page.Size = model.Page.TotalElements;
            return this.HAL(converter.Convert(model),  HttpStatusCode.OK);
        }
        // GET api/values/5
        [HttpGet("{id}", Name = "GetOffer")]
        public IActionResult Get(string id)
        {
            CustomerOrder order = collection.Find(x => x.Id.Equals(id)).First();
            if(order == null) {
                return NotFound();
            }
            return this.HAL(converter.Convert(order), HttpStatusCode.OK);
        }

        [HttpGet, Route("search/customerId/{custId?}/{sort=date}")]
        public IActionResult Get(string custId, string sort)
        {
            List<CustomerOrder> result = new List<CustomerOrder>();
            var sortBy = Builders<CustomerOrder>.Sort.Ascending(sort);
            var options = new FindOptions<CustomerOrder> { Sort = sortBy };
            result = collection.FindSync(x => x.CustomerId.Equals(custId), options).ToList();
            OrdersModel model = new OrdersModel();
            model.Orders = result;
            model.Page.TotalPages = 1;
            model.Page.TotalElements = model.Orders.Count;
            model.Page.Size = model.Page.TotalElements;
            return this.HAL(converter.Convert(model),  HttpStatusCode.OK);
        }
        [HttpPost]
        public IActionResult Create([FromBody] NewOrderResource item)
        {
            if (item == null)
            {
                return BadRequest();
            }
			Thread.Sleep(2000);
            Address address = null;
            Customer customer = null;
            Card card = null;
            List<Item> items = null;

            /* Task.Factory.ContinueWhenAll(
                new Task[] {
                    createHalAsyncTask<Address>(item.Address.AbsoluteUri)
                        .ContinueWith((task) => { address = task.Result; }),
                    createHalAsyncTask<Customer>(item.Customer.AbsoluteUri)
                        .ContinueWith((task) => { customer = task.Result; }),
                    createHalAsyncTask<Card>(item.Card.AbsoluteUri)
                        .ContinueWith((task) => { card = task.Result; }),
                    createHalAsyncTask<List<Item>>(item.Items.AbsoluteUri)
                        .ContinueWith((task) => { items = task.Result; })
                    },
                    _ => {})
                .Wait();
 */
			WebRequest cartRequest = WebRequest.Create(item.Items.AbsoluteUri);
			var cartResponse = cartRequest.GetResponseAsync().Result;
			
			StreamReader cartreader = new StreamReader(cartResponse.GetResponseStream());
			
			string responseFromServer = cartreader.ReadToEnd();
			items = JsonConvert.DeserializeObject<List<Item>>(responseFromServer);
			
			WebRequest customerRequest = WebRequest.Create(item.Customer.AbsoluteUri);
			var customerResponse = customerRequest.GetResponseAsync().Result;
			
			StreamReader customerreader = new StreamReader(customerResponse.GetResponseStream());
			
			responseFromServer = customerreader.ReadToEnd();
			customer = JsonConvert.DeserializeObject<Customer>(responseFromServer);
			
			WebRequest addressRequest = WebRequest.Create(item.Address.AbsoluteUri);
			var addressResponse = addressRequest.GetResponseAsync().Result;
			
			StreamReader addressreader = new StreamReader(addressResponse.GetResponseStream());
			
			responseFromServer = addressreader.ReadToEnd();
			address = JsonConvert.DeserializeObject<Address>(responseFromServer);
			
			WebRequest cardRequest = WebRequest.Create(item.Card.AbsoluteUri);
			var cardsResponse = cardRequest.GetResponseAsync().Result;
			
			StreamReader cardreader = new StreamReader(cardsResponse.GetResponseStream());
			
			responseFromServer = cardreader.ReadToEnd();
			card = JsonConvert.DeserializeObject<Card>(responseFromServer);
			
			Thread.Sleep(2000);
			
            PaymentResponse paymentResponse = null;

            float amount = CalculateTotal(items);
            PaymentRequest paymentRequest = new PaymentRequest() {
                Address = address,
                Card = card,
                Customer = customer,
                Amount = amount
            };
/* 
            client.PostAsync<PaymentResponse>(new HalKit.Models.Response.Link {HRef =  AppSettings.ServiceEndpoints.PaymentServiceEndpoint, IsTemplated = false}, paymentRequest)
                .ContinueWith((task) => {
                    paymentResponse = task.Result;
                })
               .Wait(); */
			   
			var data = System.Text.Encoding.ASCII.GetBytes(JsonConvert.SerializeObject(paymentRequest).ToString());
			WebRequest payRequest = WebRequest.Create(AppSettings.ServiceEndpoints.PaymentServiceEndpoint);
			payRequest.Method = "POST";
			payRequest.ContentType = "application/json";
			Stream reqStream = payRequest.GetRequestStreamAsync().Result;
			reqStream.Write(data, 0, data.Length);

			var payResponse = payRequest.GetResponseAsync().Result;
			
			StreamReader payreader = new StreamReader(payResponse.GetResponseStream());
			responseFromServer = payreader.ReadToEnd();
			paymentResponse = JsonConvert.DeserializeObject<PaymentResponse>(responseFromServer);
			
            if(!paymentResponse.Authorised) {
                return BadRequest();
            }

            string ACustomerId = customer.Id;
            Shipment Shipment = null;
                Shipment AShipment = new Shipment() {
                    Name = ACustomerId
                };
               /*  client.PostAsync<Shipment>(new HalKit.Models.Response.Link {HRef =  AppSettings.ServiceEndpoints.ShippingServiceEndpoint, IsTemplated = false}, AShipment)
                .ContinueWith((task) => {
                    Shipment = task.Result;
                })
                .Wait(); */

			var shipmentData = System.Text.Encoding.ASCII.GetBytes(JsonConvert.SerializeObject(AShipment).ToString());
			WebRequest shipmentRequest = WebRequest.Create(AppSettings.ServiceEndpoints.ShippingServiceEndpoint);
			shipmentRequest.Method = "POST";
			shipmentRequest.ContentType = "application/json";
			Stream reqShipmentStream = shipmentRequest.GetRequestStreamAsync().Result;
			reqShipmentStream.Write(shipmentData, 0, shipmentData.Length);

			var shipmentResponse = shipmentRequest.GetResponseAsync().Result;
			
			StreamReader shipmentReader = new StreamReader(shipmentResponse.GetResponseStream());
			responseFromServer = shipmentReader.ReadToEnd();
			Shipment = JsonConvert.DeserializeObject<Shipment>(responseFromServer);
			
            CustomerOrder order = new CustomerOrder() {
                CustomerId = ACustomerId,
                Address = address,
                Card  = card,
                Customer = customer,
                Items = items,
                Total = amount,
                Shipment = Shipment
            };

            collection.InsertOne(order);
            //return CreatedAtRoute("GetOffer", new { id = order.Id }, order);

            customer.Id = null;
            address.Id = null;
            card.Id = null;

            return new ObjectResult(order) {
                StatusCode = 201
            };
        }

        private float CalculateTotal(List<Item> items) {
            float amount = 0F;
            float shipping = 4.99F;
            items.ForEach(item => amount += item.Quantity * item.UnitPrice);
            amount += shipping;
            return amount;
        }

        private string ToStringNullSafe(object value) {
            return (value ?? string.Empty).ToString();
        }

        Task<T> createHalAsyncTask<T>(string link) {
            return client.GetAsync<T>(new HalKit.Models.Response.Link {HRef = link, IsTemplated = false}
                , new Dictionary<string, string> ()
                , new Dictionary<string, IEnumerable<string>>
                {
                    // it's needed to avoid
                    // org.springframework.web.HttpMediaTypeNotAcceptableException: Could not find acceptable representation
                    // from a spring based server
                    {"Accept", new[] {"application/hal+json", "application/json"}}
                });
        }
    }
}
