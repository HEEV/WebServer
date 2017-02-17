#!/bin/bash
mongod --dbpath data &
mongo &
npm start
