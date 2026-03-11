<?php

namespace App\Infrastructure\Mongo;

use MongoDB\Client;

class MongoClient
{
    public static function getClient()
    {
        return new Client("mongodb://mongo:27017");
    }
}