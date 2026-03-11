<?php

declare(strict_types=1);

namespace App\Application\Actions\User;

use Psr\Http\Message\ResponseInterface as Response;

class CreateUserAction extends UserAction
{
    /**
     * {@inheritdoc}
     */
    protected function action(): Response
    {
        $data = $this->getFormData();

        $userId = $this->userRepository->create($data);

        $this->logger->info("User created", [
            "user_id" => $userId
        ]);

        $user = $this->userRepository->findUserOfId($userId);

        return $this->respondWithData($user);
    }
}
