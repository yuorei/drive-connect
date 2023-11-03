import { Resolvers } from '../types/generated/graphql.js';
import * as mutation from './mutation/index.js';
// import * as query from './query';
import { dateScalar } from './scalar/date.js';

export const resolvers: Resolvers = {
    // Query: query,
    Mutation: mutation,
    Date: dateScalar,
};
