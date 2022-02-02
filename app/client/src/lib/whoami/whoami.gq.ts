import type * as Types from '../graphql/types.gq';

import { writable } from "svelte/store"
import { g } from '$lib/g'
import type { GFetchReturnWithErrors, GGetParameters } from '@leveluptuts/g-query'
import gql from 'graphql-tag';
export type StartWhoamiChallengeMutationVariables = Types.Exact<{
  email: Types.Scalars['String'];
}>;


export type StartWhoamiChallengeMutation = { __typename?: 'Mutation', startWhoamiChallenge?: { __typename?: 'EmptyResponse', success: boolean } | null };

export type RedeemWhoamiChallengeMutationVariables = Types.Exact<{
  email: Types.Scalars['String'];
  token: Types.Scalars['String'];
}>;


export type RedeemWhoamiChallengeMutation = { __typename?: 'Mutation', redeemWhoamiChallenge?: { __typename?: 'JwtResponse', success: boolean, accessToken: string, refreshToken: string } | null };

export type RefreshTokenMutationVariables = Types.Exact<{
  prevRefreshToken: Types.Scalars['String'];
}>;


export type RefreshTokenMutation = { __typename?: 'Mutation', refreshToken?: { __typename?: 'JwtResponse', accessToken: string, refreshToken: string, success: boolean } | null };

export type MeQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me: { __typename?: 'User', id: string, email: string } };



type FetchWrapperArgs<T> = {
	fetch: typeof fetch,
	variables?: T,
}

type SubscribeWrapperArgs<T> = {
	variables?: T,
}

interface CacheFunctionOptions {
	update?: boolean
}


export const StartWhoamiChallengeDoc = gql`
    mutation startWhoamiChallenge($email: String!) {
  startWhoamiChallenge(email: $email) {
    success
  }
}
    `;
export const RedeemWhoamiChallengeDoc = gql`
    mutation redeemWhoamiChallenge($email: String!, $token: String!) {
  redeemWhoamiChallenge(email: $email, token: $token) {
    success
    accessToken
    refreshToken
  }
}
    `;
export const RefreshTokenDoc = gql`
    mutation refreshToken($prevRefreshToken: String!) {
  refreshToken(prevRefreshToken: $prevRefreshToken) {
    accessToken
    refreshToken
    success
  }
}
    `;
export const MeDoc = gql`
    query Me {
  me {
    id
    email
  }
}
    `;

export const startWhoamiChallenge = ({ variables }: SubscribeWrapperArgs<StartWhoamiChallengeMutationVariables>):
Promise<GFetchReturnWithErrors<StartWhoamiChallengeMutation>> =>
	g.fetch<StartWhoamiChallengeMutation>({
		queries: [{ query: StartWhoamiChallengeDoc, variables }],
		fetch,
	})


export const redeemWhoamiChallenge = ({ variables }: SubscribeWrapperArgs<RedeemWhoamiChallengeMutationVariables>):
Promise<GFetchReturnWithErrors<RedeemWhoamiChallengeMutation>> =>
	g.fetch<RedeemWhoamiChallengeMutation>({
		queries: [{ query: RedeemWhoamiChallengeDoc, variables }],
		fetch,
	})


export const refreshToken = ({ variables }: SubscribeWrapperArgs<RefreshTokenMutationVariables>):
Promise<GFetchReturnWithErrors<RefreshTokenMutation>> =>
	g.fetch<RefreshTokenMutation>({
		queries: [{ query: RefreshTokenDoc, variables }],
		fetch,
	})


export const Me = writable<GFetchReturnWithErrors<MeQuery>>()

// Cached
export async function getMe({ fetch, variables }: GGetParameters<MeQueryVariables>, options?: CacheFunctionOptions) {
	const data = await g.fetch<MeQuery>({
		queries: [{ query: MeDoc, variables }],
		fetch
	})
	await Me.set({ ...data, errors: data?.errors, gQueryStatus: 'LOADED' })	
	return data
}

