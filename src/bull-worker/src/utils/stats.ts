import axios from 'axios';
import { handleError } from './handleError';
import * as Sentry from '@sentry/node';
import Redis from 'ioredis';
import { StatsResponse } from '../@types/api';

const starCitizenApi = 'https://api.starcitizen-api.com';

const redis = new Redis();

/*
RSI crowdfund
https://robertsspaceindustries.com/api/stats/getCrowdfundStat
Method: POST
Headers:
content-type: application/json
origin: https://robertsspaceindustries.com
Body:
{
  chart: "month",
  current_live: true,
  current_ptu: true,
  fans: true,
  funds: true,
  alpha_slots: true,
  fleet: true,
},
*/
/**
 * Caches game stats.
 */
export const cacheGameStats = async () => {
  // start trans
  const request = Sentry.startTransaction({
    name: 'Cache Game Stats',
    op: 'cacheGameStats',
  });

  let data: StatsResponse;

  try {
    const response = await axios({
      method: 'GET',
      url: `${starCitizenApi}/${process.env.SC_API_KEY}/v1/live/stats`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    data = response.data;
  } catch (error) {
    handleError(error);

    data = {
      data: null,
      success: 0,
      message: 'Some internal error occurred.',
    };
  } finally {
    request.finish();
  }

  const pipeline = redis.pipeline();

  if (data.data !== null) {
    pipeline
      .set('current_live', data.data.current_live)
      .set('current_ptu', data.data.current_ptu)
      .set('fans', data.data.fans)
      .set('fleet', data.data.fleet)
      .set('funds', data.data.funds);
  } else {
    pipeline
      .set('current_live', '')
      .set('current_ptu', '')
      .set('fans', '')
      .set('fleet', '')
      .set('funds', '');
  }

  pipeline.exec().catch((err) => {
    handleError(err);
  });
};
