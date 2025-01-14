/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Articles } from '../models/Articles';
import type { AuthToken } from '../models/AuthToken';
import type { Content } from '../models/Content';
import type { Contents } from '../models/Contents';
import type { PatchArticle } from '../models/PatchArticle';
import type { PatchPassword } from '../models/PatchPassword';
import type { PatchUser } from '../models/PatchUser';
import type { PostArticle } from '../models/PostArticle';
import type { PostLogin } from '../models/PostLogin';
import type { PostUsers } from '../models/PostUsers';
import type { Users } from '../models/Users';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * @param requestBody
     * @returns any Password updated
     * @throws ApiError
     */
    public static updatePassword(
        requestBody: PatchPassword,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/admin/profile/password',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
                500: `Error message`,
            },
        });
    }
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
                500: `Error message`,
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
                500: `Error message`,
            },
        });
    }
    /**
     * @returns any List of articles
     * @throws ApiError
     */
    public static getArticles(): CancelablePromise<{
        data: Articles;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/admin/articles',
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * @param requestBody
     * @returns any Article created
     * @throws ApiError
     */
    public static createArticle(
        requestBody: PostArticle,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/admin/articles',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
                500: `Error message`,
            },
        });
    }
    /**
     * @param articleIds
     * @returns any All article deleted
     * @throws ApiError
     */
    public static deleteArticles(
        articleIds: Array<number>,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/admin/articles',
            query: {
                'articleIds': articleIds,
            },
            errors: {
                500: `Error message`,
            },
        });
    }
    /**
     * @param articleId
     * @param requestBody
     * @returns any Article updated
     * @throws ApiError
     */
    public static updateArticle(
        articleId: number,
        requestBody: PatchArticle,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/admin/article/{articleId}',
            path: {
                'articleId': articleId,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Error message`,
                404: `Error message`,
                500: `Error message`,
            },
        });
    }
    /**
     * @param articleId
     * @returns any Article published
     * @throws ApiError
     */
    public static publishArticle(
        articleId: number,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/admin/article/{articleId}/publish',
            path: {
                'articleId': articleId,
            },
            errors: {
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
}
