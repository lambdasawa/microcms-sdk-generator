/* tslint:disable */
/* eslint-disable */
/**
 * microcms-sdk-generator-demo
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    UsersSns,
    UsersSnsFromJSON,
    UsersSnsFromJSONTyped,
    UsersSnsToJSON,
} from './';

/**
 * 
 * @export
 * @interface UsersPatchRequest
 */
export interface UsersPatchRequest {
    /**
     * 自己紹介
     * @type {string}
     * @memberof UsersPatchRequest
     */
    bio?: string;
    /**
     * 誕生日
     * @type {Date}
     * @memberof UsersPatchRequest
     */
    birthday?: Date;
    /**
     * メールアドレス: メールアドレスです。任意入力です。example.com である必要があります。
     * @type {string}
     * @memberof UsersPatchRequest
     */
    email?: string;
    /**
     * 
     * @type {UsersSns}
     * @memberof UsersPatchRequest
     */
    github?: UsersSns;
    /**
     * SNSを公開するか？
     * @type {boolean}
     * @memberof UsersPatchRequest
     */
    isSnsPublic?: boolean;
    /**
     * 名前
     * @type {string}
     * @memberof UsersPatchRequest
     */
    name?: string;
    /**
     * 
     * @type {UsersSns}
     * @memberof UsersPatchRequest
     */
    twitter?: UsersSns;
}

export function UsersPatchRequestFromJSON(json: any): UsersPatchRequest {
    return UsersPatchRequestFromJSONTyped(json, false);
}

export function UsersPatchRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsersPatchRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bio': !exists(json, 'bio') ? undefined : json['bio'],
        'birthday': !exists(json, 'birthday') ? undefined : (new Date(json['birthday'])),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'github': !exists(json, 'github') ? undefined : UsersSnsFromJSON(json['github']),
        'isSnsPublic': !exists(json, 'is_sns_public') ? undefined : json['is_sns_public'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'twitter': !exists(json, 'twitter') ? undefined : UsersSnsFromJSON(json['twitter']),
    };
}

export function UsersPatchRequestToJSON(value?: UsersPatchRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bio': value.bio,
        'birthday': value.birthday === undefined ? undefined : (value.birthday.toISOString().substr(0,10)),
        'email': value.email,
        'github': UsersSnsToJSON(value.github),
        'is_sns_public': value.isSnsPublic,
        'name': value.name,
        'twitter': UsersSnsToJSON(value.twitter),
    };
}


