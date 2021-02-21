import * as Sentry from '@sentry/node';
import axios from 'axios';
// import Redis from 'ioredis';
import wikijs from 'wikijs';

import { handleError } from './';
import {
  OrgResponse,
  ShipResponse,
  // StatsResponse,
  UserResponse,
  WikiArticle,
} from '../@types/api';

const wikiApiUrl = 'https://starcitizen.tools/api.php';
const starCitizenApi = 'https://api.starcitizen-api.com';

// const redis = new Redis();
const wiki = wikijs({
  apiUrl: wikiApiUrl,
});

/**
 * Searches wiki for page based on the search term
 * @param searchTerm
 */
export const getPage = async (
  searchTerm: string,
): Promise<WikiArticle | undefined> => {
  const request = Sentry.startTransaction({
    name: 'Search Page',
    op: 'searchPage',
  });

  try {
    const searchResponse = await wiki.search(searchTerm, 1);

    // if not even 1 result
    if (searchResponse.results.length < 1) {
      // end transaction
      request.finish();
      // return undefined
      return undefined;
    }

    // @ts-ignore
    const page = await wiki.page(searchResponse.results[0]);

    return {
      title: page.raw.title,
      fullURl: page.raw.fullurl,
      mainImage: await page.mainImage(),
      summary: await page.summary(),
    };
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};

/**
 * Gets a random page from the wiki
 */
export const getRandomPage = async (): Promise<
  WikiArticle | undefined
> => {
  const request = Sentry.startTransaction({
    name: 'Random Page',
    op: 'randomPage',
  });

  try {
    const random = await wiki.random(1);

    // if not even 1 result
    if (random.length < 1) {
      // end transaction
      request.finish();
      // return undefined
      return undefined;
    }

    // @ts-ignore
    const page = await wiki.page(random[0]);

    return {
      title: page.raw.title,
      fullURl: page.raw.fullurl,
      mainImage: await page.mainImage(),
      summary: await page.summary(),
    };
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};

/**
 * Gets a user based on the provided username=
 * @param username
 */
export const getUser = async (
  username: string,
): Promise<UserResponse | undefined> => {
  const request = Sentry.startTransaction({
    name: 'Get User',
    op: 'getUser',
  });

  try {
    const response = await axios({
      method: 'GET',
      url: `${starCitizenApi}/${process.env.SC_API_KEY}/v1/auto/user/${username}`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return response.data;
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};

/**
 * Gets an org based on the provided SID
 * @param SID
 */
export const getOrg = async (
  SID: string,
): Promise<OrgResponse | undefined> => {
  const request = Sentry.startTransaction({
    name: 'Get Organization',
    op: 'getOrg',
  });

  try {
    const response = await axios({
      method: 'GET',
      url: `${starCitizenApi}/${process.env.SC_API_KEY}/v1/auto/organization/${SID}`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return response.data;
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};

/**
 * WIP | Searches for ships based on the provided criteria
 * Refactor may be needed
 * @param shipName
 */
/*
export const searchShip = async (shipName: string) => {
  const request = Sentry.startTransaction({
    name: 'Get Organization',
    op: 'getOrg',
  });

  try {
    const response = await axios({
      method: 'GET',
      url: `${starCitizenApi}/${process.env.SC_API_KEY}/v1/auto/ships/?page_max=1&name=${shipName}`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return response.data;
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};
*/

/**
 * Get a ship based on the name provided
 * @param shipName
 */
export const getShip = async (
  shipName: string,
): Promise<ShipResponse | undefined> => {
  const request = Sentry.startTransaction({
    name: 'Get Ship',
    op: 'getShip',
  });

  try {
    const response = await axios({
      method: 'GET',
      url: `${starCitizenApi}/${process.env.SC_API_KEY}/v1/auto/ships/?page_max=1&name=${shipName}`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return response.data;
  } catch (error) {
    handleError(error);
  } finally {
    request.finish();
  }

  return undefined;
};
