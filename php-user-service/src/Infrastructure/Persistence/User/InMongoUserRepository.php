<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\User;

use App\Domain\User\UserRepository;
use App\Infrastructure\Mongo\MongoClient;
use MongoDB\BSON\ObjectId;
use MongoDB\BSON\UTCDateTime;

class InMongoUserRepository implements UserRepository
{
    private $collection;

    public function __construct()
    {
        $client = MongoClient::getClient();
        $this->collection = $client->gbu_db_test->users;
    }

    public function findAll(): array
    {
        $cursor = $this->collection->find([
            "deleted_at" => null
        ]);

        $users = [];

        foreach ($cursor as $doc) {
            $users[] = $this->normalize($doc);
        }

        return $users;
    }

    public function findUserOfId(string $id): array
    {
        $user = $this->collection->findOne([
            "_id" => new ObjectId($id),
            "deleted_at" => null
        ]);

        if (!$user) {
            throw new \Exception("User not found");
        }

        return $this->normalize($user);
    }

    public function create(array $data): string
    {
        $data['created_at'] = new UTCDateTime();
        $data['deleted_at'] = null;

        $result = $this->collection->insertOne($data);

        return (string) $result->getInsertedId();
    }

    public function update(string $id, array $data): bool
    {
        $data['updated_at'] = new UTCDateTime();

        $result = $this->collection->updateOne(
            [
                "_id" => new ObjectId($id),
                "deleted_at" => null
            ],
            [
                '$set' => $data
            ]
        );

        return $result->getModifiedCount() > 0;
    }

    public function delete(string $id): bool
    {
        $result = $this->collection->updateOne(
            [
                "_id" => new ObjectId($id),
                "deleted_at" => null
            ],
            [
                '$set' => [
                    "deleted_at" => new UTCDateTime()
                ]
            ]
        );

        return $result->getModifiedCount() > 0;
    }

    /**
     * Normalize Mongo document → API array
     */
    private function normalize($doc): array
    {
        $data = $doc->getArrayCopy();

        if (isset($data['_id'])) {
            $data['_id'] = (string) $data['_id'];
        }

        if (isset($data['created_at']) && $data['created_at'] instanceof UTCDateTime) {
            $data['created_at'] = $data['created_at']->toDateTime()->format(DATE_ATOM);
        }

        if (isset($data['updated_at']) && $data['updated_at'] instanceof UTCDateTime) {
            $data['updated_at'] = $data['updated_at']->toDateTime()->format(DATE_ATOM);
        }

        if (isset($data['deleted_at']) && $data['deleted_at'] instanceof UTCDateTime) {
            $data['deleted_at'] = $data['deleted_at']->toDateTime()->format(DATE_ATOM);
        }

        return $data;
    }
}