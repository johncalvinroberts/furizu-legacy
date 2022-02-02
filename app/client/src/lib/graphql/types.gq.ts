export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Date: any;
};

export type EmptyResponse = {
  __typename?: 'EmptyResponse';
  success: Scalars['Boolean'];
};

export type JwtResponse = {
  __typename?: 'JwtResponse';
  accessToken: Scalars['String'];
  refreshToken: Scalars['String'];
  success: Scalars['Boolean'];
};

export type Mutation = {
  __typename?: 'Mutation';
  redeemWhoamiChallenge?: Maybe<JwtResponse>;
  refreshToken?: Maybe<JwtResponse>;
  revokeToken?: Maybe<EmptyResponse>;
  startWhoamiChallenge?: Maybe<EmptyResponse>;
};


export type MutationRedeemWhoamiChallengeArgs = {
  email: Scalars['String'];
  token: Scalars['String'];
};


export type MutationRefreshTokenArgs = {
  prevRefreshToken: Scalars['String'];
};


export type MutationStartWhoamiChallengeArgs = {
  email: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  me: User;
};

export type User = {
  __typename?: 'User';
  createdAt: Scalars['Date'];
  email: Scalars['String'];
  id: Scalars['ID'];
  lastUpsertAt: Scalars['Date'];
};
