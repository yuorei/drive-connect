import * as grpc from '@grpc/grpc-js';
import { UserServiceClient } from '../../proto/user_grpc_pb';
import * as user from '../../proto/user_pb';
import { MutationResolvers, User } from '../../types/generated/graphql';
export const createUser: MutationResolvers['createUser'] = async (
    parent,
    args,
    context,
    info
) => {
    const client = new UserServiceClient(
        'localhost:50051',
        grpc.credentials.createInsecure()
    );
    const request = new user.User();
    request.setName(args.input.name);
    request.setEmail(args.input.email);
    request.setPassword(args.input.password);
    request.setIsDriver(args.input.is_driver);

    let userGrpc = await new Promise<user.User>((resolve, reject) => {
        client.createUser(request, (err, response) => {
            if (err) {
                return reject(err);
            }

            return resolve(response);
        });
    });

    const userResponse: User = {
        id: userGrpc.getId(),
        name: userGrpc.getName(),
        email: userGrpc.getEmail(),
        is_driver: userGrpc.getIsDriver(),
        created_at: userGrpc.getCreatedAt().toDate(),
        updated_at: userGrpc.getUpdatedAt().toDate(),
    };

    return userResponse;
}

