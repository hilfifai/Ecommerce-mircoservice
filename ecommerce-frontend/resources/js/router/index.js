import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import ProductDetail from '../views/ProductDetail.vue';
import Login from '../views/Login.vue';
import Cart from '../views/Cart.vue';
import Register from '../views/Register.vue'; 
import Checkout from '../views/Checkout.vue';

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/login', name: 'Login', component: Login },
    { path: '/product/:id', name: 'ProductDetail', component: ProductDetail, props: true },
    { path: '/cart', name: 'Cart', component: Cart },
    {
        path: '/register', 
        name: 'Register',
        component: Register
    }, { 
        path: '/checkout',
        name: 'Checkout', 
        component: Checkout 
    },
 
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
      
        redirect: '/' 
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;