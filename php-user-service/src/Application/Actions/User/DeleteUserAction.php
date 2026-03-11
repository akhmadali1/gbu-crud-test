<?php

declare(strict_types=1);

namespace App\Application\Actions\User;

use Psr\Http\Message\ResponseInterface as Response;

class DeleteUserAction extends UserAction
{
    /**
     * {@inheritdoc}
     */
    protected function action(): Response
    {
        $id = (string) $this->resolveArg('id');

        $this->userRepository->delete($id);

        $this->logger->info("User deleted", [
            "user_id" => $id
        ]);

        return $this->respondWithData([
            "status" => "deleted",
            "user_id" => $id
        ]);
    }
}