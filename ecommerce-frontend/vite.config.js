import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
    plugins: [
        laravel({
            input: [
                'resources/sass/app.scss',
                'resources/css/app.css',
                'resources/js/app.js',
            ],
            refresh: true,
        }),
        vue({
            template: {
                transformAssetUrls: {
                    base: null,
                    includeAbsolute: false,
                },
            },
        }),
    ],
    server: {
        proxy: {
        '/api': {
            target: 'http://localhost:8080',
            changeOrigin: true,
            rewrite: path => path.replace(/^\/api/, '')
        }
        }
    },
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm-bundler.js',
        },
    },
});
