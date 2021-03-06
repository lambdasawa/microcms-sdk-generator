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
/**
 * 
 * @export
 * @interface UsersSns
 */
export interface UsersSns {
    /**
     * 
     * @type {string}
     * @memberof UsersSns
     */
    fieldId?: UsersSnsFieldIdEnum;
    /**
     * URL
     * @type {string}
     * @memberof UsersSns
     */
    url?: string;
    /**
     * ユーザID
     * @type {string}
     * @memberof UsersSns
     */
    userId?: string;
}

/**
* @export
* @enum {string}
*/
export enum UsersSnsFieldIdEnum {
    Sns = 'sns'
}

export function UsersSnsFromJSON(json: any): UsersSns {
    return UsersSnsFromJSONTyped(json, false);
}

export function UsersSnsFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsersSns {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fieldId': !exists(json, 'fieldId') ? undefined : json['fieldId'],
        'url': !exists(json, 'url') ? undefined : json['url'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
    };
}

export function UsersSnsToJSON(value?: UsersSns | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fieldId': value.fieldId,
        'url': value.url,
        'userId': value.userId,
    };
}


