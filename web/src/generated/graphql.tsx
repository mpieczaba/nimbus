import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
  Upload: any;
};



export type AuthPayload = {
  __typename?: 'AuthPayload';
  /** JWT authorization token */
  token: Scalars['String'];
  /** Signed in user */
  user: User;
};

export type File = {
  __typename?: 'File';
  /** Unique id */
  id: Scalars['ID'];
  /** File name */
  name: Scalars['String'];
  /** File [mime type](https://www.iana.org/assignments/media-types/media-types.xhtml) */
  mimeType: Scalars['String'];
  /** File extension */
  extension: Scalars['String'];
  /** File size */
  size: Scalars['Int'];
  /** File URL */
  url: Scalars['String'];
  /** Tags file was tagged with */
  tags?: Maybe<FileTagConnection>;
  /** File collaborators */
  collaborators?: Maybe<FileCollaboratorConnection>;
  /** Create time */
  createdAt: Scalars['Time'];
  /** Update time */
  updatedAt: Scalars['Time'];
};


export type FileTagsArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  name?: Maybe<Scalars['String']>;
};


export type FileCollaboratorsArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  username?: Maybe<Scalars['String']>;
  permission?: Maybe<FilePermission>;
};

export type FileCollaboratorConnection = {
  __typename?: 'FileCollaboratorConnection';
  edges?: Maybe<Array<Maybe<FileCollaboratorEdge>>>;
  nodes?: Maybe<Array<Maybe<User>>>;
  pageInfo: PageInfo;
};

export type FileCollaboratorEdge = {
  __typename?: 'FileCollaboratorEdge';
  cursor: Scalars['String'];
  node: User;
  permission: FilePermission;
};

export type FileCollaboratorInput = {
  /** File id */
  fileId: Scalars['ID'];
  /** Collaborator user id */
  collaboratorId: Scalars['ID'];
  /** File permission */
  permission: FilePermission;
};

export type FileConnection = {
  __typename?: 'FileConnection';
  edges?: Maybe<Array<Maybe<FileEdge>>>;
  nodes?: Maybe<Array<Maybe<File>>>;
  pageInfo: PageInfo;
};

export type FileEdge = {
  __typename?: 'FileEdge';
  cursor: Scalars['String'];
  node: File;
};

export type FileInput = {
  /** File name (optional) */
  name?: Maybe<Scalars['String']>;
  /** File */
  file: Scalars['Upload'];
};

export enum FilePermission {
  Admin = 'ADMIN',
  Maintain = 'MAINTAIN',
  Write = 'WRITE',
  Read = 'READ'
}

export type FileTagConnection = {
  __typename?: 'FileTagConnection';
  edges?: Maybe<Array<Maybe<FileTagEdge>>>;
  nodes?: Maybe<Array<Maybe<Tag>>>;
  pageInfo: PageInfo;
};

export type FileTagEdge = {
  __typename?: 'FileTagEdge';
  cursor: Scalars['String'];
  node: Tag;
};

export type FileTagsInput = {
  /** File id */
  fileId: Scalars['ID'];
  /** Tag names */
  tagNames: Array<Scalars['String']>;
};

export type FileUpdateInput = {
  /** File name */
  name?: Maybe<Scalars['String']>;
  /** File */
  file?: Maybe<Scalars['Upload']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  /** Sign in user with username and password */
  login: AuthPayload;
  /** Sign up user with UserInput */
  createUser?: Maybe<User>;
  /** Update currently authenticated user with UserUpdateInput or user with given id */
  updateUser?: Maybe<User>;
  /** Delete currently authenticated user or user with with id */
  deleteUser?: Maybe<User>;
  /** Create new File with FileInput */
  createFile?: Maybe<File>;
  /** Update file with id and FileUpdateInput */
  updateFile?: Maybe<File>;
  /** Delete file with id */
  deleteFile?: Maybe<File>;
  /** Add tags to the file with FileTagsInput */
  addTagsToFile?: Maybe<File>;
  /** Remove tags from file with FileTagsInput */
  removeTagsFromFile?: Maybe<File>;
  /** Add file collaborator with FileCollaboratorInput */
  addCollaboratorToFile?: Maybe<File>;
  /** Remove file collaborator from the file with file id and collaborator id */
  removeCollaboratorFromFile?: Maybe<File>;
};


export type MutationLoginArgs = {
  username: Scalars['String'];
  password: Scalars['String'];
};


export type MutationCreateUserArgs = {
  input: UserInput;
};


export type MutationUpdateUserArgs = {
  id?: Maybe<Scalars['ID']>;
  input: UserUpdateInput;
};


export type MutationDeleteUserArgs = {
  id?: Maybe<Scalars['ID']>;
};


export type MutationCreateFileArgs = {
  input: FileInput;
};


export type MutationUpdateFileArgs = {
  id: Scalars['ID'];
  input: FileUpdateInput;
};


export type MutationDeleteFileArgs = {
  id: Scalars['ID'];
};


export type MutationAddTagsToFileArgs = {
  input: FileTagsInput;
};


export type MutationRemoveTagsFromFileArgs = {
  input: FileTagsInput;
};


export type MutationAddCollaboratorToFileArgs = {
  input: FileCollaboratorInput;
};


export type MutationRemoveCollaboratorFromFileArgs = {
  fileId: Scalars['ID'];
  collaboratorId: Scalars['ID'];
};

export type PageInfo = {
  __typename?: 'PageInfo';
  hasNextPage: Scalars['Boolean'];
  hasPreviousPage: Scalars['Boolean'];
  startCursor?: Maybe<Scalars['String']>;
  endCursor?: Maybe<Scalars['String']>;
};

export type Query = {
  __typename?: 'Query';
  /** Get user by id or currently authenticated user if no id is specified */
  user?: Maybe<User>;
  /** Get all users */
  users?: Maybe<UserConnection>;
  /** Get file by id */
  file?: Maybe<File>;
  /** Get all files */
  files?: Maybe<FileConnection>;
  /** Get tag by name */
  tag?: Maybe<Tag>;
  /** Get all tags */
  tags?: Maybe<TagConnection>;
};


export type QueryUserArgs = {
  id?: Maybe<Scalars['ID']>;
};


export type QueryUsersArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  username?: Maybe<Scalars['String']>;
};


export type QueryFileArgs = {
  id: Scalars['ID'];
};


export type QueryFilesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  name?: Maybe<Scalars['String']>;
  permission?: Maybe<FilePermission>;
  tags?: Maybe<Array<Scalars['String']>>;
};


export type QueryTagArgs = {
  name: Scalars['String'];
};


export type QueryTagsArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  name?: Maybe<Scalars['String']>;
};

export type Tag = {
  __typename?: 'Tag';
  /** Unique name */
  name: Scalars['String'];
  /** Files tagged with the tag */
  files?: Maybe<FileConnection>;
  /** Create time */
  createdAt: Scalars['Time'];
  /** Update time */
  updatedAt: Scalars['Time'];
};


export type TagFilesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  name?: Maybe<Scalars['String']>;
  permission?: Maybe<FilePermission>;
  tags?: Maybe<Array<Scalars['String']>>;
};

export type TagConnection = {
  __typename?: 'TagConnection';
  edges?: Maybe<Array<Maybe<TagEdge>>>;
  nodes?: Maybe<Array<Maybe<Tag>>>;
  pageInfo: PageInfo;
};

export type TagEdge = {
  __typename?: 'TagEdge';
  cursor: Scalars['String'];
  node: Tag;
};



export type User = {
  __typename?: 'User';
  /** Unique id */
  id: Scalars['ID'];
  /** Unique username */
  username: Scalars['String'];
  /** User kind */
  kind: UserKind;
  /** Create time */
  createdAt: Scalars['Time'];
  /** Update time */
  updatedAt: Scalars['Time'];
};

export type UserConnection = {
  __typename?: 'UserConnection';
  edges?: Maybe<Array<Maybe<UserEdge>>>;
  nodes?: Maybe<Array<Maybe<User>>>;
  pageInfo: PageInfo;
};

export type UserEdge = {
  __typename?: 'UserEdge';
  cursor: Scalars['String'];
  node: User;
};

export type UserInput = {
  /** Unique username */
  username: Scalars['String'];
  /** User password */
  password: Scalars['String'];
};

export enum UserKind {
  Admin = 'ADMIN',
  User = 'USER',
  Banned = 'BANNED'
}

export type UserUpdateInput = {
  /** Unique username */
  username?: Maybe<Scalars['String']>;
  /** User password */
  password?: Maybe<Scalars['String']>;
  /** User kind - requires admin privileges */
  kind?: Maybe<UserKind>;
};

export type LoginMutationVariables = Exact<{
  username: Scalars['String'];
  password: Scalars['String'];
}>;


export type LoginMutation = (
  { __typename?: 'Mutation' }
  & { login: (
    { __typename?: 'AuthPayload' }
    & Pick<AuthPayload, 'token'>
    & { user: (
      { __typename?: 'User' }
      & Pick<User, 'username'>
    ) }
  ) }
);

export type FilesQueryVariables = Exact<{
  after?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
  name?: Maybe<Scalars['String']>;
  tags?: Maybe<Array<Scalars['String']> | Scalars['String']>;
}>;


export type FilesQuery = (
  { __typename?: 'Query' }
  & { files?: Maybe<(
    { __typename?: 'FileConnection' }
    & { edges?: Maybe<Array<Maybe<(
      { __typename?: 'FileEdge' }
      & Pick<FileEdge, 'cursor'>
      & { node: (
        { __typename?: 'File' }
        & Pick<File, 'id' | 'name' | 'size' | 'url' | 'updatedAt'>
      ) }
    )>>>, pageInfo: (
      { __typename?: 'PageInfo' }
      & Pick<PageInfo, 'hasNextPage' | 'endCursor'>
    ) }
  )> }
);

export type TagsQueryVariables = Exact<{
  name?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int']>;
}>;


export type TagsQuery = (
  { __typename?: 'Query' }
  & { tags?: Maybe<(
    { __typename?: 'TagConnection' }
    & { edges?: Maybe<Array<Maybe<(
      { __typename?: 'TagEdge' }
      & { node: (
        { __typename?: 'Tag' }
        & Pick<Tag, 'name'>
      ) }
    )>>> }
  )> }
);


export const LoginDocument = gql`
    mutation Login($username: String!, $password: String!) {
  login(username: $username, password: $password) {
    token
    user {
      username
    }
  }
}
    `;
export type LoginMutationFn = Apollo.MutationFunction<LoginMutation, LoginMutationVariables>;

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginMutation, { data, loading, error }] = useLoginMutation({
 *   variables: {
 *      username: // value for 'username'
 *      password: // value for 'password'
 *   },
 * });
 */
export function useLoginMutation(baseOptions?: Apollo.MutationHookOptions<LoginMutation, LoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<LoginMutation, LoginMutationVariables>(LoginDocument, options);
      }
export type LoginMutationHookResult = ReturnType<typeof useLoginMutation>;
export type LoginMutationResult = Apollo.MutationResult<LoginMutation>;
export type LoginMutationOptions = Apollo.BaseMutationOptions<LoginMutation, LoginMutationVariables>;
export const FilesDocument = gql`
    query Files($after: String, $first: Int, $name: String, $tags: [String!]) {
  files(after: $after, first: $first, name: $name, tags: $tags) {
    edges {
      cursor
      node {
        id
        name
        size
        url
        updatedAt
      }
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
}
    `;

/**
 * __useFilesQuery__
 *
 * To run a query within a React component, call `useFilesQuery` and pass it any options that fit your needs.
 * When your component renders, `useFilesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useFilesQuery({
 *   variables: {
 *      after: // value for 'after'
 *      first: // value for 'first'
 *      name: // value for 'name'
 *      tags: // value for 'tags'
 *   },
 * });
 */
export function useFilesQuery(baseOptions?: Apollo.QueryHookOptions<FilesQuery, FilesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<FilesQuery, FilesQueryVariables>(FilesDocument, options);
      }
export function useFilesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<FilesQuery, FilesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<FilesQuery, FilesQueryVariables>(FilesDocument, options);
        }
export type FilesQueryHookResult = ReturnType<typeof useFilesQuery>;
export type FilesLazyQueryHookResult = ReturnType<typeof useFilesLazyQuery>;
export type FilesQueryResult = Apollo.QueryResult<FilesQuery, FilesQueryVariables>;
export const TagsDocument = gql`
    query Tags($name: String, $first: Int) {
  tags(name: $name, first: $first) {
    edges {
      node {
        name
      }
    }
  }
}
    `;

/**
 * __useTagsQuery__
 *
 * To run a query within a React component, call `useTagsQuery` and pass it any options that fit your needs.
 * When your component renders, `useTagsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useTagsQuery({
 *   variables: {
 *      name: // value for 'name'
 *      first: // value for 'first'
 *   },
 * });
 */
export function useTagsQuery(baseOptions?: Apollo.QueryHookOptions<TagsQuery, TagsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<TagsQuery, TagsQueryVariables>(TagsDocument, options);
      }
export function useTagsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<TagsQuery, TagsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<TagsQuery, TagsQueryVariables>(TagsDocument, options);
        }
export type TagsQueryHookResult = ReturnType<typeof useTagsQuery>;
export type TagsLazyQueryHookResult = ReturnType<typeof useTagsLazyQuery>;
export type TagsQueryResult = Apollo.QueryResult<TagsQuery, TagsQueryVariables>;