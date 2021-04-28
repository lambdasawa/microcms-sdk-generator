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
 * @interface SettingPatchRequest
 */
export interface SettingPatchRequest {
    /**
     * 1ページあたりの記事数
     * @type {number}
     * @memberof SettingPatchRequest
     */
    articlesPerPage?: number;
    /**
     * 特集記事
     * @type {Array<string>}
     * @memberof SettingPatchRequest
     */
    featuredArticles?: Array<string>;
}

export function SettingPatchRequestFromJSON(json: any): SettingPatchRequest {
    return SettingPatchRequestFromJSONTyped(json, false);
}

export function SettingPatchRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SettingPatchRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'articlesPerPage': !exists(json, 'articlesPerPage') ? undefined : json['articlesPerPage'],
        'featuredArticles': !exists(json, 'featuredArticles') ? undefined : json['featuredArticles'],
    };
}

export function SettingPatchRequestToJSON(value?: SettingPatchRequest | null): any {
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


