/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 * 
 *   http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License. 
 */

"use client";

import { Card } from "@mui/material";
import React, { ReactNode } from "react";
import { ClusterCreation } from "./formCreation";
import { faCirclePlus } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

interface CreateCardProps {
  children: ReactNode;
}

export const CreateCard: React.FC<CreateCardProps> = ({ children }) => {
    return (
        <Card
            variant="outlined"
            sx={{
                width: "300px",
                height: "200px",
                padding: "16px",
                margin: "16px",
                borderRadius: "16px",
                transition: "transform 0.1s, box-shadow 0.3s",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                "&:hover": {
                    transform: "scale(1.01)",
                    boxShadow: "0 4px 8px rgba(0, 0, 0, 0.2)",
                },
                cursor: "pointer",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
            }}
        >
            {children}
        </Card>
    );
};

export const AddClusterCardProps = ({ namespace }: { namespace: string }) => {
    return (
        <>
            <FontAwesomeIcon
                icon={faCirclePlus}
                size="4x"
                style={{
                    color: "#e0e0e0",
                    marginBottom: "8px",
                    transition: "color 0.2s",
                }}
            />
            <div className="mt-4">
                <ClusterCreation position="card" namespace={namespace} />
            </div>
        </>
    );
};