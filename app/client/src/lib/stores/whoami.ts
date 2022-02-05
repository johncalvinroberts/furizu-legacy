import { writable } from 'svelte/store';
import { browser } from '$app/env';
import decodeJwt from 'jwt-decode';
import { TOKEN, REFRESH_TOKEN } from '../constants';

type Whoami = {
	isLoading: boolean;
	isLoggedIn: boolean;
	accessTokenExpires?: number;
	accessToken?: string | undefined;
	refreshToken?: string | undefined;
	email?: string | undefined;
	userId?: string | undefined;
};

type AccessTokenPayload = {
	email: string;
	expire: number;
	iat: number;
	userId: string;
};

type TokenSet = {
	accessToken: string | undefined;
	refreshToken: string | undefined;
};

const noopValue = {
	isLoggedIn: false,
	isLoading: true,
};

const computeWhoamiState = ({ accessToken, refreshToken }: TokenSet): Whoami => {
	let isLoggedIn = false;
	let accessTokenPayload: AccessTokenPayload;
	let accessTokenExpires;
	if (refreshToken && accessToken) {
		accessTokenPayload = decodeJwt<AccessTokenPayload>(accessToken);
		accessTokenExpires = accessTokenPayload.expire;
		const refreshTokenPayload = decodeJwt<AccessTokenPayload>(accessToken);
		const now = new Date().valueOf();
		const isRefreshTokenExpired = refreshTokenPayload.expire * 1000 < now;
		if (isRefreshTokenExpired) {
			localStorage?.removeItem(TOKEN);
			localStorage?.removeItem(REFRESH_TOKEN);
		}
		isLoggedIn = !isRefreshTokenExpired;
	}
	return {
		isLoading: false,
		accessToken,
		refreshToken,
		userId: accessTokenPayload?.userId,
		email: accessTokenPayload?.email,
		isLoggedIn,
		accessTokenExpires,
	};
};

const getWhoamiFromLocalStorage = () => {
	const accessToken = localStorage?.getItem(TOKEN);
	const refreshToken = localStorage?.getItem(REFRESH_TOKEN);
	return computeWhoamiState({ refreshToken, accessToken });
};

const initialValue = browser ? getWhoamiFromLocalStorage() : noopValue;
export const whoami = writable<Whoami>(initialValue);

export const setWhoamiState = (tokens: TokenSet) => {
	localStorage.setItem(TOKEN, tokens.accessToken);
	localStorage.setItem(REFRESH_TOKEN, tokens.refreshToken);
	const nextState = computeWhoamiState(tokens);
	whoami.set(nextState);
};

export const logout = () => {
	localStorage?.removeItem(TOKEN);
	localStorage?.removeItem(REFRESH_TOKEN);
	const nextState = getWhoamiFromLocalStorage();
	whoami.set(nextState);
};
