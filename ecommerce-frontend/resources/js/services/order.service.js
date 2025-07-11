import { backendApi } from './api';

class OrderService {
    createOrder(orderPayload) {
      
        // { items: [{ product_id: 1, quantity: 2 }, ...] }
       
        return backendApi.post('/orders', orderPayload);
    }
}

export default new OrderService();