package works.weave.socks.orders.resources;

import java.net.URI;
import java.util.List;

import org.hibernate.validator.constraints.URL;

import works.weave.socks.orders.entities.Item;

public class NewOrderResource {
    @URL
    public URI customer;

    @URL
    public URI address;

    @URL
    public URI card;

    public List<Item> items;

    public String customerId;
}
