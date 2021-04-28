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
 * @interface SettingCreateRequest
 */
export interface SettingCreateRequest {
    /**
     * 1ページあたりの記事数
     * @type {number}
     * @memberof SettingCreateRequest
     */
    articlesPerPage: number;
    /**
     * 特集記事
     * @type {Array<string>}
     * @memberof SettingCreateRequest
     */
    featuredArticles: Array<string>;
}

export function SettingCreateRequestFromJSON(json: any): SettingCreateRequest {
    return SettingCreateRequestFromJSONTyped(json, false);
}

export function SettingCreateRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SettingCreateRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'articlesPerPage': json['articlesPerPage'],
        'featuredArticles': json['featuredArticles'],
    };
}

export function SettingCreateRequestToJSON(value?: SettingCreateRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'articlesPerPage': value.articlesPerPage,
        'featuredArticles': value.featuredArticles,
    };
}


