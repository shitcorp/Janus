import { Sequelize, Model, DataTypes } from 'sequelize';

// start sequelize
const sequelize = new Sequelize('sqlite::memory:');

class Prefix extends Model {}

Prefix.init(
  {
    guildID: DataTypes.STRING,
    prefix: DataTypes.STRING,
  },
  { sequelize, modelName: 'prefix' },
);

export const getPrefix = (guildID: string) => {};

export const setPrefix = async (
  guildID: string,
  prefix: string,
) => {
  await sequelize.sync();

  const jane = await Prefix.create({
    guildID: guildID,
    prefix: prefix,
  });

  console.log(jane.toJSON());
};
