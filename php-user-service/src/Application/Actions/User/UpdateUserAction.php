<?php

declare(strict_types=1);

namespace App\Application\Actions\User;

use Psr\Http\Message\ResponseInterface as Response;

class UpdateUserAction extends UserAction
{
    /**
     * {@inheritdoc}
     */
    protected function action(): Response
    {
        $id = (string) $this->resolveArg('id');

        $data = $this->getFormData();

        $this->userRepository->update($id, $data);

        $this->logger->info("User updated", [
            "user_id" => $id
        ]);

        $user = $this->userRepository->findUserOfId($id);

        return $this->respondWithData($user);
    }
}