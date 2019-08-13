package works.weave.socks.orders.values;

import works.weave.socks.orders.entities.Address;
import works.weave.socks.orders.entities.Card;
import works.weave.socks.orders.entities.Customer;

public class PaymentRequest {

    private float amount;

    // For jackson
    public PaymentRequest() {
    }

    public PaymentRequest(Address address, Card card, Customer customer, float amount) {
        this.amount = amount;
    }

    public PaymentRequest(float amount) {
        this.amount = amount;
    }

    @Override
    public String toString() {
        return "PaymentRequest{" +
                '}';
    }

    public float getAmount() {
        return amount;
    }

    public void setAmount(float amount) {
        this.amount = amount;
    }
}
