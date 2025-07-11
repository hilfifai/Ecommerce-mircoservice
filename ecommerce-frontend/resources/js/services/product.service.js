import { directusApi } from './api';

const PRODUCT_FIELDS = 'id,name,description,price,stock'; // Tentukan field yang dibutuhkan

class ProductService {
    getProducts() {
        return directusApi.get('/items/products', {
            params: {
                fields: PRODUCT_FIELDS,
                filter: {
                    status: { _eq: 'published' }
                }
            }
        });
    }

    getProductById(id) {
        return directusApi.get(`/items/products/${id}`, {
            params: {
                fields: PRODUCT_FIELDS
            }
        });
    }
}

export default new ProductService();