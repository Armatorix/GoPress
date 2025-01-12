/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AuthToken } from '../models/AuthToken';
import type { Content } from '../models/Content';
import type { Contents } from '../models/Contents';
import type { PatchContent } from '../models/PatchContent';
import type { PatchUser } from '../models/PatchUser';
import type { PostContent } from '../models/PostContent';
import type { PostLogin } from '../models/PostLogin';
import type { PostUsers } from '../models/PostUsers';
import type { Users } from '../models/Users';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * @returns Users List of users
     * @throws ApiError
     */
    public static getUsers(): CancelablePromise<Users> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/admin/users',
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * Create a new user
     * @param requestBody
     * @returns any User created
     * @throws ApiError
     */
    public static postAdminUsers(
        requestBody: PostUsers,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/admin/users',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
            },
        });
    }
    /**
     * @param userIds
     * @returns any All users deleted
     * @throws ApiError
     */
    public static deleteUsers(
        userIds: Array<string>,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/admin/users',
            query: {
                'userIds': userIds,
            },
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * @param userId
     * @param requestBody
     * @returns any User updated
     * @throws ApiError
     */
    public static updateUser(
        userId: string,
        requestBody: PatchUser,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/admin/users/{userId}',
            path: {
                'userId': userId,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
                500: `Error message`,
            },
        });
    }
    /**
     * @returns Contents List of contents
     * @throws ApiError
     */
    public static getContents(): CancelablePromise<Contents> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/admin/contents',
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * Create new content
     * @param requestBody
     * @returns any Content created
     * @throws ApiError
     */
    public static postAdminContents(
        requestBody: PostContent,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/admin/contents',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
            },
        });
    }
    /**
     * Delete all content
     * @returns any All content deleted
     * @throws ApiError
     */
    public static deleteAdminContents(): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/admin/contents',
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * Update content
     * @param contentId
     * @param requestBody
     * @returns any Content updated
     * @throws ApiError
     */
    public static patchAdminContent(
        contentId: string,
        requestBody: PatchContent,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/admin/content/{contentId}',
            path: {
                'contentId': contentId,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
                404: `Error message`,
            },
        });
    }
    /**
     * @param contentId
     * @returns Content Content
     * @throws ApiError
     */
    public static getContent(
        contentId: string,
    ): CancelablePromise<Content> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/admin/content/{contentId}',
            path: {
                'contentId': contentId,
            },
            errors: {
                404: `Error message`,
                500: `Error message`,
            },
        });
    }
    /**
     * Login
     * @param requestBody
     * @returns AuthToken Login response
     * @throws ApiError
     */
    public static postAdminLogin(
        requestBody: PostLogin,
    ): CancelablePromise<AuthToken> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/admin/login',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
            },
        });
    }
}
