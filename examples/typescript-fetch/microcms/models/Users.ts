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
    UsersIcon,
    UsersIconFromJSON,
    UsersIconFromJSONTyped,
    UsersIconToJSON,
    UsersSns,
    UsersSnsFromJSON,
    UsersSnsFromJSONTyped,
    UsersSnsToJSON,
} from './';

/**
 * 
 * @export
 * @interface Users
 */
export interface Users {
    /**
     * 自己紹介
     * @type {string}
     * @memberof Users
     */
    bio: string;
    /**
     * 誕生日
     * @type {Date}
     * @memberof Users
     */
    birthday: Date;
    /**
     * 
     * @type {Date}
     * @memberof Users
     */
    createdAt: Date;
    /**
     * メールアドレス: メールアドレスです。任意入力です。example.com である必要があります。
     * @type {string}
     * @memberof Users
     */
    email: string;
    /**
     * 
     * @type {UsersSns}
     * @memberof Users
     */
    github?: UsersSns;
    /**
     * 
     * @type {UsersIcon}
     * @memberof Users
     */
    icon?: UsersIcon;
    /**
     * 
     * @type {string}
     * @memberof Users
     */
    id: string;
    /**
     * SNSを公開するか？
     * @type {boolean}
     * @memberof Users
     */
    isSnsPublic: boolean;
    /**
     * 名前
     * @type {string}
     * @memberof Users
     */
    name: string;
    /**
     * 
     * @type {Date}
     * @memberof Users
     */
    publishedAt: Date;
    /**
     * 
     * @type {Date}
     * @memberof Users
     */
    revisedAt: Date;
    /**
     * 
     * @type {UsersSns}
     * @memberof Users
     */
    twitter?: UsersSns;
    /**
     * 
     * @type {Date}
     * @memberof Users
     */
    updatedAt: Date;
}

export function UsersFromJSON(json: any): Users {
    return UsersFromJSONTyped(json, false);
}

export function UsersFromJSONTyped(json: any, ignoreDiscriminator: boolean): Users {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bio': json['bio'],
        'birthday': (new Date(json['birthday'])),
        'createdAt': (new Date(json['createdAt'])),
        'email': json['email'],
        'github': !exists(json, 'github') ? undefined : UsersSnsFromJSON(json['github']),
        'icon': !exists(json, 'icon') ? undefined : UsersIconFromJSON(json['icon']),
        'id': json['id'],
        'isSnsPublic': json['is_sns_public'],
        'name': json['name'],
        'publishedAt': (new Date(json['publishedAt'])),
        'revisedAt': (new Date(json['revisedAt'])),
        'twitter': !exists(json, 'twitter') ? undefined : UsersSnsFromJSON(json['twitter']),
        'updatedAt': (new Date(json['updatedAt'])),
    };
}

export function UsersToJSON(value?: Users | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bio': value.bio,
        'birthday': (value.birthday.toISOString().substr(0,10)),
        'createdAt': (value.createdAt.toISOString()),
        'email': value.email,
        'github': UsersSnsToJSON(value.github),
        'icon': UsersIconToJSON(value.icon),
        'id': value.id,
        'is_sns_public': value.isSnsPublic,
        'name': value.name,
        'publishedAt': (value.publishedAt.toISOString()),
        'revisedAt': (value.revisedAt.toISOString()),
        'twitter': UsersSnsToJSON(value.twitter),
        'updatedAt': (value.updatedAt.toISOString()),
    };
}


