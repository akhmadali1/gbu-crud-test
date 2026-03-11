<?php

declare(strict_types=1);

namespace App\Domain\User;

use JsonSerializable;

class User implements JsonSerializable
{
    private ?string $id;

    public string $user_name;

    public string $first_name;

    public ?string $last_name = null;

    public \DateTime $created_at;

    public ?\DateTime $updated_at = null;

    public ?\DateTime $deleted_at = null;

    public function __construct(?string $id, string $user_name, string $first_name, string $last_name)
    {
        $this->_id = $id;
        $this->user_name = ($user_name);
        $this->first_name = ($first_name);
        $this->last_name = ($last_name ?? null);

        $this->created_at = new \DateTime();
    }

    public function getId(): ?int
    {
        return $this->_id;
    }

    public function getUsername(): string
    {
        return $this->user_name;
    }

    public function getFirstName(): string
    {
        return $this->first_name;
    }

    public function getLastName(): string
    {
        return $this->last_name;
    }

    #[\ReturnTypeWillChange]
    public function jsonSerialize(): array
    {
        return [
            "_id" => $this->_id,
            "user_name"  => $this->user_name,
            "first_name" => $this->first_name,
            "last_name"  => $this->last_name,
            "created_at" => $this->created_at,
            "updated_at" => $this->updated_at,
            "deleted_at" => $this->deleted_at,
        ];
    }
}
