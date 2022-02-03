import { writable } from 'svelte/store';
import { browser } from '$app/env';
import decodeJwt from 'jwt-decode';
import { TOKEN, REFRESH_TOKEN } from '../constants';

type Whoami = {
	isLoggedIn: boolean;
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
};

const computeWhoamiState = ({ accessToken, refreshToken }: TokenSet): Whoami => {
	let isLoggedIn = false;
	let accessTokenPayload: AccessTokenPayload;
	if (refreshToken && accessToken) {
		accessTokenPayload = decodeJwt<AccessTokenPayload>(accessToken);
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
		accessToken,
		refreshToken,
		userId: accessTokenPayload?.userId,
		email: accessTokenPayload?.email,
		isLoggedIn,
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
	const nextState = computeWhoamiState(tokens);
	whoami.set(nextState);
};
