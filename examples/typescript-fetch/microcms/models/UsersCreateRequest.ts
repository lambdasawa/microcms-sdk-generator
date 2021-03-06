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
 * @interface UsersCreateRequest
 */
export interface UsersCreateRequest {
    /**
     * 自己紹介
     * @type {string}
     * @memberof UsersCreateRequest
     */
    bio: string;
    /**
     * 誕生日
     * @type {Date}
     * @memberof UsersCreateRequest
     */
    birthday: Date;
    /**
     * メールアドレス: メールアドレスです。任意入力です。example.com である必要があります。
     * @type {string}
     * @memberof UsersCreateRequest
     */
    email: string;
    /**
     * 
     * @type {UsersSns}
     * @memberof UsersCreateRequest
     */
    github?: UsersSns;
    /**
     * SNSを公開するか？
     * @type {boolean}
     * @memberof UsersCreateRequest
     */
    isSnsPublic: boolean;
    /**
     * 名前
     * @type {string}
     * @memberof UsersCreateRequest
     */
    name: string;
    /**
     * その他のSNSのリスト
     * @type {Array<UsersSns>}
     * @memberof UsersCreateRequest
     */
    otherSNSList?: Array<UsersSns>;
    /**
     * 
     * @type {UsersSns}
     * @memberof UsersCreateRequest
     */
    twitter?: UsersSns;
}

export function UsersCreateRequestFromJSON(json: any): UsersCreateRequest {
    return UsersCreateRequestFromJSONTyped(json, false);
}

export function UsersCreateRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsersCreateRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bio': json['bio'],
        'birthday': (new Date(json['birthday'])),
        'email': json['email'],
        'github': !exists(json, 'github') ? undefined : UsersSnsFromJSON(json['github']),
        'isSnsPublic': json['is_sns_public'],
        'name': json['name'],
        'otherSNSList': !exists(json, 'otherSNSList') ? undefined : ((json['otherSNSList'] as Array<any>).map(UsersSnsFromJSON)),
        'twitter': !exists(json, 'twitter') ? undefined : UsersSnsFromJSON(json['twitter']),
    };
}

export function UsersCreateRequestToJSON(value?: UsersCreateRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bio': value.bio,
        'birthday': (value.birthday.toISOString().substr(0,10)),
        'email': value.email,
        'github': UsersSnsToJSON(value.github),
        'is_sns_public': value.isSnsPublic,
        'name': value.name,
        'otherSNSList': value.otherSNSList === undefined ? undefined : ((value.otherSNSList as Array<any>).map(UsersSnsToJSON)),
        'twitter': UsersSnsToJSON(value.twitter),
    };
}


