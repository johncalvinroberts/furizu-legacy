import { GFetch } from '@leveluptuts/g-query';

const API_BASE_URL = process.env.NODE_ENV === 'development' ? 'http://localhost:4000/' : '/';

export const g = new GFetch({
	path: `${API_BASE_URL}query`, //whatever your api url is here
});
